package main

/*
#cgo CFLAGS: -DPNG_DEBUG=1 -I./ffmpeg/include -I./hls
#cgo LDFLAGS: -L./ffmpeg/lib -llibavformat -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
#include <hls.c>
*/
import "C"

func main() {
	inFilename := "rtsp://183.59.168.27/PLTV/88888905/224/3221227272/10000100000000060000000001030757_0.smil?icip=88888888"
	outFilename := "D:/Env/nginx/html/hls/ffmpeg/test.m3u8"
	C.tohls(C.CString(inFilename), C.CString(outFilename))
}
