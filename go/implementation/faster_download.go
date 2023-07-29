package implementation

import "context"

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
func (d *Downloader) Download(url string, fileOut string) error {
	return d.DownloadWithContext(context.Background(), url, fileOut)
}

// Like Download, but with context
//
// You can use context to cancel the download
func (d *Downloader) DownloadWithContext(ctx context.Context, url string, fileOut string) error {
	// TODO: Implement this function
	return nil
}
