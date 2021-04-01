package main

import (
	"github.com/snowlyg/go_ffmpeg"
)

func main() {
	hls := &go_ffmpeg.Hls{
		InFilename:  "rtsp://admin:P@ssw0rd@10.0.0.10:554/Streaming/Channels/102",
		OutFilename: "./hls_files",
		HlsTime:     "2",
		HlsListSize: "5",
	}

	err := hls.ToHls()
	if err != nil {
		println(err)
	}
	select {}
}
