package main

import (
	"os"

	"github.com/snowlyg/go_ffmpeg/src"
)

func main() {
	go hls(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
	select {}
}

func hls(in, out, ht, hl string) {
	if ht == "" {
		ht = "2"
	}

	if hl == "" {
		hl = "5"
	}

	hls := &src.Hls{
		InFilename:  in,
		OutFilename: out,
		HlsTime:     ht,
		HlsListSize: hl,
	}

	err := hls.ToHls()
	if err != nil {
		println(err)
	}
}
