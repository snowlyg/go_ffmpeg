package src

/*
#cgo CFLAGS: -I/usr/include/ffmpeg
#cgo LDFLAGS: -L/usr/lib64/ -lswscale  -lavcodec -lavformat -lavutil -lswresample -lavdevice -lavfilter
#include <hls/hls.c>
*/
import "C"
import (
	"fmt"
)

func (h *Hls) ToHls() error {
	err := CreateFile(h.OutFilename)
	if err != nil {
		return fmt.Errorf("create file %w", err)
	}
	outFilename := h.OutFilename + "/out.m3u8"
	C.to_hls(C.CString(h.InFilename), C.CString(outFilename), C.CString(h.RtspTransport.String()), C.CString(h.HlsTime), C.CString(h.HlsListSize))

	return nil
}
