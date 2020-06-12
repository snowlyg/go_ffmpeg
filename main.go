package main

/*
#cgo CFLAGS: -DPNG_DEBUG=1 -I./ffmpeg/include -I./hls
#cgo LDFLAGS: -L${SRCDIR}/ffmpeg/lib -llibavformat -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
#include <hls.c>
*/
import "C"

func main() {
	C.tohls()
}
