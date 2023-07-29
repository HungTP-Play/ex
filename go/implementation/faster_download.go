package implementation

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

// ChunkSizeDeterminer is a function that determines the chunk size based on the file size
type ChunkSizeDeterminer func(fileSize int64) int64

// PartDeterminer is a function that determines the number of parts to download concurrently
type PartDeterminer func(fileSize int64) int64

// DownloaderConfig is a struct that contains configuration for the Downloader
type DownloaderConfig struct {
	// MaxRetries is the maximum number of retries for a file to download
	// If MaxRetries is -1, it will retry forever
	//
	// If one of chunk download fails, it will retry entire file download
	//
	// Default is 5
	MaxRetries int

	// MaxChunkSize is the maximum size of a chunk in bytes
	// If MaxChunkSize is -1, it will download the entire file in one chunk
	//
	// Default is 1MB
	MaxChunkSize int

	// The Concurrency is the number of chunks to download concurrently
	//
	// Default: -1 (download all chunks concurrently)
	Concurrency int

	// ChunkSizeDeterminer is a function that determines the chunk size based on the file size
	// If ChunkSizeDeterminer is nil, it will use the default chunk size
	//
	// Default: 1MB
	// The ChunkSize after calling ChunkSizeDeterminer will be capped by MaxChunkSize -> min(ChunkSize, MaxChunkSize)
	//
	// You should only use ChunkSizeDeterminer either PartDeterminer, not both
	ChunkSizeDeterminer ChunkSizeDeterminer

	// PartDeterminer is a function that determines the number of parts to download concurrently
	// If PartDeterminer is nil, it will use the default number of parts -> ceil(fileSize / chunkSize)
	//
	// If the number of parts is greater than Concurrency, some parts will be downloaded concurrently, and some parts will be downloaded sequentially
	//
	// Example:
	// Concurrency = 3
	// PartDeterminer = func(fileSize int64) int64 {
	// 	return 2
	// }
	// The first 2 parts will be downloaded concurrently, and the last part will be downloaded sequentially
	//
	// You should only use PartDeterminer either ChunkSizeDeterminer, not both
	//
	// By default,  Downloader will use the default PartDeterminer
	PartDeterminer PartDeterminer

	// The ChunkSize is the size of a chunk in bytes
	chunkSize int64
}

// File Size Unit Enum
const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

var defaultPartDeterminer = func(fileSize int64) int64 {
	if fileSize < 1*MB {
		return 1
	}

	if fileSize < 10*MB {
		return 2
	}

	if fileSize < 100*MB {
		return 16
	}

	return 32
}

var defaultConfig = DownloaderConfig{
	MaxRetries:          5,
	MaxChunkSize:        1 * MB,
	Concurrency:         -1,
	ChunkSizeDeterminer: nil,
	PartDeterminer:      defaultPartDeterminer,
}

type Downloader struct {
	config DownloaderConfig
}

func NewDownloader(options ...Option) *Downloader {
	downloader := Downloader{config: defaultConfig}
	for _, option := range options {
		option(&downloader)
	}
	return &downloader
}

func NewDownloaderWithConfig(config DownloaderConfig) *Downloader {
	return &Downloader{config: config}
}

type Option func(d *Downloader)

func WithMaxRetries(maxRetries int) Option {
	return func(d *Downloader) {
		d.config.MaxRetries = maxRetries
	}
}

func WithMaxChunkSize(maxChunkSize int) Option {
	return func(d *Downloader) {
		d.config.MaxChunkSize = maxChunkSize
	}
}

func WithConcurrency(concurrency int) Option {
	return func(d *Downloader) {
		d.config.Concurrency = concurrency
	}
}

func WithChunkSizeDeterminer(chunkSizeDeterminer ChunkSizeDeterminer) Option {
	return func(d *Downloader) {
		d.config.ChunkSizeDeterminer = chunkSizeDeterminer
	}
}

func WithPartDeterminer(partDeterminer PartDeterminer) Option {
	return func(d *Downloader) {
		d.config.PartDeterminer = partDeterminer
	}
}

// Download downloads a file from url and save it to fileOut
func (d *Downloader) Download(url string, fileOut string) (int64, error) {
	return d.DownloadWithContext(context.Background(), url, fileOut)
}

type FileWriteAt struct {
	file *os.File
}

func (f *FileWriteAt) WriteAt(p []byte, off int64) (int, error) {
	return f.file.WriteAt(p, off)
}

// Like Download, but with context
//
// You can use context to cancel the download
func (d *Downloader) DownloadWithContext(ctx context.Context, url string, fileOut string) (int64, error) {
	// Open the file
	// If the file doesn't exist, create it, or append to the file
	file, err := os.OpenFile(fileOut, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	w := &FileWriteAt{file: file}
	impl := downloader{w: w, cfg: d.config, ctx: ctx, url: url, out: fileOut}

	fileSize, err := impl.getFileSize()
	if err != nil {
		return 0, err
	}

	impl.totalBytes = fileSize
	if d.config.Concurrency == -1 {
		d.config.Concurrency = int(defaultPartDeterminer(fileSize))
		chunkSize := math.Min(float64(fileSize)/float64(d.config.Concurrency), float64(d.config.MaxChunkSize))
		d.config.chunkSize = int64(chunkSize)
	}

	impl.cfg.Concurrency = d.config.Concurrency
	impl.cfg.chunkSize = d.config.chunkSize
	return impl.download()
}

// downloader is the implementation structure used internally by Downloader.
type downloader struct {
	ctx context.Context
	cfg DownloaderConfig

	url string
	out string

	w  io.WriterAt
	wg sync.WaitGroup
	m  sync.Mutex

	pos        int64
	totalBytes int64
	written    int64
	err        error
}

// dlchunk represents a single chunk of data to write by the worker routine.
// This structure also implements an io.SectionReader style interface for
// io.WriterAt, effectively making it an io.SectionWriter (which does not
// exist).
type dlchunk struct {
	w     io.WriterAt
	start int64
	size  int64
	cur   int64

	// specifies the byte range the chunk should be downloaded with.
	withRange string
}

// Write wraps io.WriterAt for the dlchunk, writing from the dlchunk's start
// position to its end (or EOF).
//
// If a range is specified on the dlchunk the size will be ignored when writing.
// as the total size may not of be known ahead of time.
func (c *dlchunk) Write(p []byte) (n int, err error) {
	if c.cur >= c.size && len(c.withRange) == 0 {
		return 0, io.EOF
	}

	n, err = c.w.WriteAt(p, c.start+c.cur)
	c.cur += int64(n)

	return
}

// ByteRange returns a HTTP Byte-Range header value that should be used by the
// client to request the chunk's range.
func (c *dlchunk) ByteRange() string {
	if len(c.withRange) != 0 {
		return c.withRange
	}

	return fmt.Sprintf("bytes=%d-%d", c.start, c.start+c.size-1)
}

type errReadingBody struct {
	err error
}

func (e *errReadingBody) Error() string {
	return fmt.Sprintf("failed to read part body: %v", e.err)
}

func (e *errReadingBody) Unwrap() error {
	return e.err
}

func (d *downloader) getHead() (*http.Response, error) {
	req, err := http.NewRequest("HEAD", d.url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req.WithContext(d.ctx))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *downloader) getFileSize() (int64, error) {
	resp, err := d.getHead()
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp.ContentLength, nil
}

// Download file from url and save it to fileOut using GETs
func (d *downloader) download() (int64, error) {

	d.getChunk()

	if total := d.getTotalBytes(); total >= 0 {
		// Spin up workers
		ch := make(chan dlchunk, d.cfg.Concurrency)

		for i := 0; i < d.cfg.Concurrency; i++ {
			d.wg.Add(1)
			go d.downloadPart(ch)
		}

		// Assign work
		for d.getErr() == nil {
			if d.pos >= total {
				break // We're finished queuing chunks
			}

			// Queue the next range of bytes to read.
			ch <- dlchunk{w: d.w, start: d.pos, size: d.cfg.chunkSize}
			d.pos += d.cfg.chunkSize
		}

		// Wait for completion
		close(ch)
		d.wg.Wait()
	} else {
		// Checking if we read anything new
		for d.err == nil {
			d.getChunk()
		}

		// We expect a 416 error letting us know we are done downloading the
		// total bytes. Since we do not know the content's length, this will
		// keep grabbing chunks of data until the range of bytes specified in
		// the request is out of range of the content. Once, this happens, a
		// 416 should occur.
		var responseError interface {
			HTTPStatusCode() int
		}
		if errors.As(d.err, &responseError) {
			if responseError.HTTPStatusCode() == http.StatusRequestedRangeNotSatisfiable {
				d.err = nil
			}
		}
	}
	// Return error
	return d.written, d.err
}

// downloadPart is an individual goroutine worker reading from the ch channel
// and performing a GetObject request on the data with a given byte range.
//
// If this is the first worker, this operation also resolves the total number
// of bytes to be read so that the worker manager knows when it is finished.
func (d *downloader) downloadPart(ch chan dlchunk) {
	defer d.wg.Done()
	for {
		chunk, ok := <-ch
		if !ok {
			break
		}
		if d.getErr() != nil {
			// Drain the channel if there is an error, to prevent deadlocking
			// of download producer.
			continue
		}

		if err := d.downloadChunk(chunk); err != nil {
			d.setErr(err)
		}
	}
}

func (d *downloader) getTotalBytes() int64 {
	d.m.Lock()
	defer d.m.Unlock()

	return d.totalBytes
}

// getChunk grabs a chunk of data from the body.
// Not thread safe. Should only used when grabbing data on a single thread.
func (d *downloader) getChunk() {
	if d.getErr() != nil {
		return
	}

	chunk := dlchunk{w: d.w, start: d.pos, size: d.cfg.chunkSize}
	d.pos += d.cfg.chunkSize

	if err := d.downloadChunk(chunk); err != nil {
		d.setErr(err)
	}
}

func (d *downloader) incrWritten(n int64) {
	d.m.Lock()
	defer d.m.Unlock()

	d.written += n
}

func (d *downloader) downloadChunk(chunk dlchunk) error {
	rangeHeaderValue := chunk.ByteRange()

	retry := 0
	var err error
	var n int64

	for ; retry < d.cfg.MaxRetries; retry++ {
		n, err = d.tryDownloadChunk(&chunk, rangeHeaderValue)
		if err == nil {
			break
		}

		// Check if the returned error is an errReadingBody.
		// If err is errReadingBody this indicates that an error
		// occurred while copying the http response body.
		// If this occurs we unwrap the err to set the underlying error
		// and attempt any remaining retries.
		if bodyErr, ok := err.(*errReadingBody); ok {
			err = bodyErr.Unwrap()
		} else {
			return err
		}

		chunk.cur = 0
	}

	d.incrWritten(n)
	return err
}

func (d *downloader) tryDownloadChunk(w io.Writer, rangeHeaderValue string) (int64, error) {
	cleanup := func() {}
	req, err := http.NewRequest("GET", d.url, nil)
	if err != nil {
		return 0, err
	}
	defer cleanup()

	req.Header.Set("Range", rangeHeaderValue)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	d.setTotalBytes(resp) // Set total if not yet set.
	defer resp.Body.Close()

	var src io.Reader = resp.Body

	n, err := io.Copy(w, src)
	if err != nil {
		return n, &errReadingBody{err: err}
	}

	return n, nil

}

// getContentRange extracts the total bytes from the Content-Range header.
func (d *downloader) getContentRange(resp *http.Response) (string, error) {
	contentRange := resp.Header.Get("Content-Range")
	if len(contentRange) == 0 {
		return "", fmt.Errorf("missing Content-Range header")
	}

	return contentRange, nil
}

// setTotalBytes is a thread-safe setter for setting the total byte status.
// Will extract the object's total bytes from the Content-Range if the file
// will be chunked, or Content-Length. Content-Length is used when the response
// does not include a Content-Range. Meaning the object was not chunked. This
// occurs when the full file fits within the PartSize directive.
func (d *downloader) setTotalBytes(resp *http.Response) {
	d.m.Lock()
	defer d.m.Unlock()

	if d.totalBytes >= 0 {
		return
	}

	contentRange, err := d.getContentRange(resp)
	if err != nil {
		// ContentRange is nil when the full file contents is provided, and
		// is not chunked. Use ContentLength instead.
		if resp.ContentLength > 0 {
			d.totalBytes = resp.ContentLength
			return
		}
	} else {
		parts := strings.Split(contentRange, "/")

		total := int64(-1)
		var err error
		// Checking for whether or not a numbered total exists
		// If one does not exist, we will assume the total to be -1, undefined,
		// and sequentially download each chunk until hitting a 416 error
		totalStr := parts[len(parts)-1]
		if totalStr != "*" {
			total, err = strconv.ParseInt(totalStr, 10, 64)
			if err != nil {
				d.err = err
				return
			}
		}

		d.totalBytes = total
	}
}

// getErr is a thread-safe getter for the error object
func (d *downloader) getErr() error {
	d.m.Lock()
	defer d.m.Unlock()

	return d.err
}

// setErr is a thread-safe setter for the error object
func (d *downloader) setErr(e error) {
	d.m.Lock()
	defer d.m.Unlock()

	d.err = e
}
