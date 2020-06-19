# ffmpegTest
尝试直接调用 ffmpeg 动态库函数,生成 hls 文件. 目前支持 windows 环境， 其他环境需要自行安装相关依赖库

##### 参考资料
[https://github.com/leandromoreira/ffmpeg-libav-tutorial#learn-ffmpeg-libav-the-hard-way](https://github.com/leandromoreira/ffmpeg-libav-tutorial#learn-ffmpeg-libav-the-hard-way)

##### 编译
gcc -c -o hls.o hls.c -I../ffmpeg/include -L../ffmpeg/lib -llibavformat -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale

##### 使用方法
- 复制所有 dll 到项目根目录下
- example
```go
package main

import "github.com/snowlyg/ffmpegTest"

func main() {
	inFilename := "rtsp://183.59.168.27/PLTV/88888905/224/3221227272/10000100000000060000000001030757_0.smil?icip=88888888"
    outFilename := "D:/Env/nginx/html/hls/ffmpeg"	
    ffmpegTest.ToHls(inFilename, outFilename)
}   

```