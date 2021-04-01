package main

import (
	"log"
	"net/http"
	"time"

	"github.com/snowlyg/go_ffmpeg"
	"github.com/snowlyg/go_ffmpeg/cmd/routers"
)

func main() {
	// 初始化路由
	err := routers.Init()
	if err != nil {
		return
	}

	go hls()
	go startHTTP()

	select {}
}

func hls() {
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
}

func startHTTP() {
	httpServer := http.Server{
		Addr:              ":10008",
		Handler:           routers.Router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	link := "http://localhost:10008"
	log.Println("http server start -->", link)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("start http server error", err)
		}
		log.Println("http server start")
	}()
}
