
import fs from 'fs/promises';
import fetch from 'node-fetch';

function getChunkSize(size: number): number {
    if (size < 1 * 1024 * 1024) return size;
    if (size < 10 * 1024 * 1024) return 1 * 1024 * 1024;
    if (size < 100 * 1024 * 1024) return 5 * 1024 * 1024;
    if (size < 1024 * 1024 * 1024) return 10 * 1024 * 1024;
    return 20 * 1024 * 1024;
}

function getNumberOfPart(size: number): number {
    if (size < 1 * 1024 * 1024) return 1;
    if (size < 10 * 1024 * 1024) return 4;
    if (size < 100 * 1024 * 1024) return 8;
    if (size < 1024 * 1024 * 1024) return 16;
    return 32;
}

async function getFileSize(url:string): Promise<number> {
    const response = await fetch(url, {
        method: 'HEAD'
    });
    const size = response.headers.get('Content-Length');
    return Number(size);
}

async function downloadPart(url:string, outPath: string, start: number, size: number): Promise<void> {
    const response = await fetch(url, {
        headers: {
            Range: `bytes=${start}-${start + size - 1}`
        }
    });
    const buffer = await response.arrayBuffer();
    await fs.writeFile(outPath, Buffer.from(buffer));
}

async function dowloadFile(url:string, outPath: string): Promise<void>{
    const fileSize = await getFileSize(url);
    const chunkSize = getChunkSize(fileSize);
    const numParts = Math.ceil(fileSize / chunkSize);

    // For using fixed number of parts
    // const numParts = getNumberOfPart(fileSize);
    // const chunkSize = Math.ceil(fileSize / numParts);

    const promises = [];

    for (let i = 0; i < numParts; i++) {
        const start = i * chunkSize;
        const end = Math.min((i + 1) * chunkSize, fileSize);
        const partSize = end - start;
        const partPath = `${outPath}.${i}`;
        promises.push(downloadPart(url, partPath, start, partSize));
    }

    await Promise.all(promises);
}