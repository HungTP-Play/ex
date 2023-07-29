package main

import (
	"github.com/HungTP-Play/ex/go/implementation"
)

func main() {
	fileOut := "book_notes.png"
	url := "https://files.catbox.moe/67bo9v.png"
	downloader := implementation.Downloader{}
	downloader.Download(url, fileOut)
}
