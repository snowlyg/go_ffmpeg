package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/snowlyg/go_ffmpeg/example/hls_server/routers"
	"github.com/snowlyg/go_ffmpeg/src"
)

func main() {
	// 初始化路由
	err := routers.Init()
	if err != nil {
		return
	}

	go hls(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
	go startHTTP()

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
