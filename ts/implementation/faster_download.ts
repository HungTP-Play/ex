/**
 * function to determine the chunk size to download.
 * 
 * @param size The size of the file.
 */
export type ChunkSizeDeterminer = (size: number) => number;

/**
 * Function to determine the number of parts to download.
 * 
 * @param size The size of the file.
 */
export type PartDeterminer = (size: number) => number;

/**
 * Config for the downloader.
 */
export type DownloaderConfig = {
    /**
     * The maximum number of retries for a chunk.
     */
    maxRetries?: number;

    /**
     * The maximum number of chunks to download at once.
     */
    maxChunkSize?: number;

    /**
     * The number of chunks to download at once.
     */
    concurrency: number;

    /**
     * The function to determine the number of parts to download.
     */
    chunkSizeDeterminer?: ChunkSizeDeterminer;

    /**
     * The function to determine the chunk size to download.
     */
    partDeterminer?: PartDeterminer;

    /**
     * The actual chunk size to download.
     */
    chunkSize: number;
};


export interface WriteAt {
    WriteAt(bytes: Uint8Array, offset: number): Promise<number>;
}

export enum SizeUnit {
    B = 1,
    KB = 1024,
    MB = 1024 * 1024,
    GB = 1024 * 1024 * 1024
}

const defaultPartDeterminer: PartDeterminer = (size: number) => {
    if (size < SizeUnit.MB) {
        return 1;
    }

    if (size < 10 * SizeUnit.GB) {
        return 4;
    }

    if (size < 100 * SizeUnit.GB) {
        return 16;
    }

    return 32;
}

const defaultConfig: DownloaderConfig = {
    maxRetries: 5,
    maxChunkSize: SizeUnit.MB,
    concurrency: 4,
    partDeterminer: defaultPartDeterminer,
    chunkSize: 0
}

export class downloadChunk {
    start: number;
    end: number;
    size: number;
    cur: number;
    writer: WriteAt;

    constructor(start: number, end: number, size: number) {
        this.start = start;
        this.end = end;
        this.size = size;
        this.cur = start;
    }

    async write(bytes: Uint8Array): Promise<number> {
        if (this.cur + bytes.length > this.end) {
            throw new Error('Invalid write');
        }

        const byteWrote =  await this.writer.WriteAt(bytes, this.start+this.cur);
        this.cur += byteWrote;
        return byteWrote;
    }

    byteRage(): string {
        return `bytes=${this.start}-${this.end}`;
    }
}

class downloader {
    constructor(private config: DownloaderConfig) { }
}

export class Downloader {
    private downloader: downloader;
    constructor(private config: DownloaderConfig) {
        this.config = Object.assign({}, defaultConfig, config);
        this.downloader = new downloader(this.config);
    }
}