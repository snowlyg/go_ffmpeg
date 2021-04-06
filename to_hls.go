/**
* useage case
* package main
*
* import "github.com/snowlyg/go_ffmpeg"
*
* func main() {
*
* 	hls := go_ffmpeg.Hls{
* 		InFilename:    "rtsp://www.mym9.com/101065?from=2019-06-28/01:12:13",
* 		OutFilename:   "./hls_files",
* 		RtspTransport: go_ffmpeg.TCP,
* 	}
*
* 	hls.ToHls()
* }
**/

package go_ffmpeg

/*
#cgo CFLAGS: -I/usr/include/ffmpeg
#cgo LDFLAGS: -L/usr/lib64/ -lswscale  -lavcodec -lavformat -lavutil -lswresample -lavdevice -lavfilter
#include <hls/hls.c>
*/
import "C"
import (
	"fmt"
	"os"
)

type RtspTransport int

const (
	TCP RtspTransport = iota
	UDP
)

func (r RtspTransport) String() string {
	switch r {
	case TCP:
		return "tcp"
	case UDP:
		return "udp"
	default:
		return "tcp"
	}
}

type Hls struct {
	InFilename    string
	OutFilename   string
	HlsTime       string
	HlsListSize   string
	RtspTransport RtspTransport
}

//	养鸡rtsp回放：rtsp://www.mym9.com/101065?from=2019-06-28/01:12:13
//	rtmp://58.200.131.2:1935/livetv/hunantv
//	inFilename := "rtsp://183.59.168.27/PLTV/88888905/224/3221227272/10000100000000060000000001030757_0.smil?icip=88888888"
//	outFilename := "D:/Env/nginx/html/hls/ffmpeg/test.m3u8"
//	rtspTransport := "tcp"

func (h *Hls) ToHls() error {
	err := CreateFile(h.OutFilename)
	if err != nil {
		return fmt.Errorf("create file %w", err)
	}
	outFilename := h.OutFilename + "/out.m3u8"
	C.to_hls(C.CString(h.InFilename), C.CString(outFilename), C.CString(h.RtspTransport.String()), C.CString(h.HlsTime), C.CString(h.HlsListSize))

	return nil
}

// CreateFile 调用os.MkdirAll递归创建文件夹
func CreateFile(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// IsExist  判断所给路径文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
