# ffmpegTest
尝试直接调用 ffmpeg 动态库函数,生成 hls 文件. 目前支持 windows 环境， 其他环境需要自行安装相关依赖库

##### 参考资料
[https://github.com/leandromoreira/ffmpeg-libav-tutorial#learn-ffmpeg-libav-the-hard-way](https://github.com/leandromoreira/ffmpeg-libav-tutorial#learn-ffmpeg-libav-the-hard-way)

##### 编译
gcc -c -o hls.o hls.c -I../ffmpeg/include -L../ffmpeg/lib -llibavformat -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
