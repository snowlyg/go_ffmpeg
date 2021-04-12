package src

import (
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
