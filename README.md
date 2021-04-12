# go_ffmpeg
本库尝试使用 `go` 直接调用 ffmpeg 动态库函数,生成 hls 文件.

#### 程序依赖
- gcc 因为需要调用 ffmpeg 的依赖库，需要通过 `cgo` 调用。
- ffmpeg(4.0+) 或者 ffmepg 依赖库(4.0+)

##### 使用
```go
package main

import "github.com/snowlyg/go_ffmpeg"

func main() {
	hls := go_ffmpeg.Hls{
		InFilename: "rtsp://www.mym9.com/101065?from=2019-06-28/01:12:13",
		OutFilename: "./hls_files",
	}

	hls.ToHls()
}   

```
更多使用示例，参考 [example](./example) 。

#### 下载 `ffmpeg` 及相关依赖


- `windows` 环境: 到 [https://github.com/BtbN/FFmpeg-Builds/releases](https://github.com/BtbN/FFmpeg-Builds/releases) 下载 `ffmpeg` 依赖库文件。 
- 

- `mac` 环境: 
```shell
 brew install ffmpeg
```

- `ubuntu` 环境: 
```shell
sudo apt install -y libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev
``` 

- `centos7` 环境: 
- [https://linuxize.com/post/how-to-install-ffmpeg-on-centos-7/](https://linuxize.com/post/how-to-install-ffmpeg-on-centos-7)

#### 配置 `cgo`
- 修改 [chls.go](src/chls.go) 文件，将 `#cgo CFLAGS: -I` 和 `cgo LDFLAGS: -L` 修改为前面依赖库的安装地址。
- `windows` 环境需要将 `dll` 文件复制到程序的执行目录


![cctv9.png](cctv9.png)


##### 参考资料
- [https://github.com/leandromoreira/ffmpeg-libav-tutorial#learn-ffmpeg-libav-the-hard-way](https://github.com/leandromoreira/ffmpeg-libav-tutorial#learn-ffmpeg-libav-the-hard-way)

- [https://linuxize.com/post/how-to-install-ffmpeg-on-centos-7](https://linuxize.com/post/how-to-install-ffmpeg-on-centos-7)

- [https://www.cnblogs.com/wanggang123/p/10302023.html](https://www.cnblogs.com/wanggang123/p/10302023.html)

- [http://www.chungen90.com/?news_34/](http://www.chungen90.com/?news_34/)
