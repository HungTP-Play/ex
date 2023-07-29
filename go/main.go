package main

import (
	"github.com/HungTP-Play/ex/go/implementation"
)

func main() {
	fileOut := "book_notes.jpg"
	url := "https://files.catbox.moe/em97nz.jpg"
	downloader := implementation.NewDownloader()
	downloader.Download(url, fileOut)
}
