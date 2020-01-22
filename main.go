package main

import (
	"os"
	"grabgravatar"
)

func main() {
	if len(os.Args) <= 2 {
		panic("need email to look for and image with. Example: ./Grabgravatar foo@example.com 1024")
	}
	mail := os.Args[1]
	size := os.Args[2]

	if err := grabgravatar.DownloadFile(mail, size); err != nil {
		panic(err)
	}
}