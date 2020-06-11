package main

/*
#cgo CFLAGS: -DPNG_DEBUG=1 -I./ffmpeg/include
#cgo LDFLAGS: -L${SRCDIR}/ffmpeg/lib -llibavformat -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
#include <libavutil/opt.h>
#include <libavdevice/avdevice.h>
#include <stdio.h>
#include <libswscale/swscale.h>

int hls()
{
    AVFormatContext *pFormatCtx;
    int             i, index;
    AVCodecContext  *pCodecCtx;
    AVCodec         *pCodec;
    AVFrame *pFrame,*pFrameYUV;
    uint8_t *out_buffer;
    AVPacket *packet;
    int y_size;
    int ret, got_picture;
    struct SwsContext *img_convert_ctx;

    // input & output init
    char *filepath = "rtsp://183.59.168.27/PLTV/88888905/224/3221227041/10000100000000060000000000657998_0.smil?icip=88888888";
    FILE *fp_yuv=fopen("output.yuv","wb+");
    FILE *fp_h264=fopen("output.h264","wb+");

    // 注册编码解码器
   av_register_all();

    // 初始化网络组件
    avformat_network_init();

    // 初始化 format
    pFormatCtx = avformat_alloc_context();
	AVDictionary *options = NULL;
	av_dict_set(&options,"rtsp_transport","tcp", 0);
    //打开视频流
    if(avformat_open_input(&pFormatCtx,filepath,NULL, &options)!=0){
        printf("Couldn't open input stream.\n");
        return -1;
    }

    // 寻找视频流
    if(avformat_find_stream_info(pFormatCtx,NULL)<0){
        printf("Couldn't find stream information.\n");
        return -1;
    }
    index = -1;
    for(i=0; i < pFormatCtx->nb_streams; i++){
        if(pFormatCtx->streams[i]->codecpar->codec_type==AVMEDIA_TYPE_VIDEO){
            index = i;
            break;
        }
    }

    if(index==-1){
        printf("Didn't find a video stream.\n");
        return -1;
    }

    // 获取源视频流的 codec
    pCodecCtx=pFormatCtx->streams[index]->codecpar;

    //查找解码器
    pCodec=avcodec_find_decoder(pCodecCtx->codec_id);
    if(pCodec==NULL){
        printf("Codec not found decodec.\n");
        return -1;
    }

    //打开解码器
    if(avcodec_open2(pCodecCtx, pCodec,NULL)<0){
        printf("Could not open codec.\n");
        return -1;
    }

    // 初始化AVFrame
    pFrame=av_frame_alloc();
    pFrameYUV=av_frame_alloc();

    out_buffer=(uint8_t *)av_malloc(av_samples_get_buffer_size(AV_PIX_FMT_YUV420P, pCodecCtx->width, pCodecCtx->height, AV_SAMPLE_FMT_S16, 1));
    //avpicture_fill((AVPicture *)pFrameYUV, out_buffer, AV_PIX_FMT_YUV420P, pCodecCtx->width, pCodecCtx->height);
	av_samples_fill_arrays(pFrameYUV->data, pFrameYUV->linesize, out_buffer, AV_PIX_FMT_YUV420P, pCodecCtx->width, pCodecCtx->height, 1);

    packet=(AVPacket *)av_malloc(sizeof(AVPacket));

    //Output Info-----------------------------
    printf("--------------- File Information ----------------\n");
    av_dump_format(pFormatCtx,0,filepath,0);
    printf("-------------------------------------------------\n");

    // sws 初始化，设置源pixfmt 和目标pixfmt
    img_convert_ctx = sws_getContext(pCodecCtx->width, pCodecCtx->height, pCodecCtx->pix_fmt,
        pCodecCtx->width, pCodecCtx->height, AV_PIX_FMT_YUV420P, SWS_BICUBIC, NULL, NULL, NULL);

    while(av_read_frame(pFormatCtx, packet)>=0){//读取一帧压缩数据
        if(packet->stream_index == index){

            fwrite(packet->data,1,packet->size,fp_h264); //把H264数据写入fp_h264文件

            //ret = avcodec_decode_video2(pCodecCtx, pFrame, &got_picture, packet);//解码一帧压缩数据
			ret = avcodec_send_packet(pCodecCtx, packet);
			got_picture = avcodec_receive_frame(pCodecCtx, pFrame);
            if(ret < 0){
                printf("Decode Error.\n");
                return -1;
            }
            if(got_picture){
                //PixelFormat 转化，转成yuv420
                sws_scale(img_convert_ctx, (const uint8_t* const*)pFrame->data, pFrame->linesize, 0, pCodecCtx->height,pFrameYUV->data, pFrameYUV->linesize);

                y_size = pCodecCtx->width * pCodecCtx->height;
                fwrite(pFrameYUV->data[0],y_size,1,fp_yuv);    //Y
                fwrite(pFrameYUV->data[1],y_size/4,1,fp_yuv);  //U
                fwrite(pFrameYUV->data[2],y_size/4,1,fp_yuv);  //V
                printf("Succeed to decode 1 frame!\n");

            }
        }
        av_packet_unref(packet);
    }

    sws_freeContext(img_convert_ctx);

    //关闭文件，释放内存
    fclose(fp_yuv);
    fclose(fp_h264);

    av_frame_free(&pFrameYUV);
    av_frame_free(&pFrame);
    avcodec_close(pCodecCtx);
    avformat_close_input(&pFormatCtx);

    return 0;
}
*/
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println(C.hls())
}
