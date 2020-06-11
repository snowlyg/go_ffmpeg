package main

/*
#cgo CFLAGS: -DPNG_DEBUG=1 -I./ffmpeg/include
#cgo LDFLAGS: -L${SRCDIR}/ffmpeg/lib -llibavformat -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
//#include <stdafx.h>
#include <libavcodec/avcodec.h>
#include <libavformat/avformat.h>
#include <libswscale/swscale.h>
#include <libavutil/imgutils.h>
#include <stdio.h>

void SaveFrame(AVFrame *pFrame, int width, int height, int iFrame) {
    FILE *pFile;
    char szFilename[32];
    int  y;

    // Open file
    sprintf_s(szFilename, "frame%d.ppm", iFrame);
    fopen_s(&pFile, szFilename, "wb");
    if (pFile == NULL)
        return;

    // Write header
    fprintf(pFile, "P6\n%d %d\n255\n", width, height);

    // Write pixel data
    for (y = 0; y < height; y++)
        fwrite(pFrame->data[0] + y*pFrame->linesize[0], 1, width * 3, pFile);

    // Close file
    fclose(pFile);
}

void hls()
{
   // Initalizing these to NULL prevents segfaults!
    AVFormatContext   *pFormatCtx = NULL;
    int               i, videoStream;
    AVCodecParameters    *pCodecpar = NULL;
    AVCodecContext    *pCodecCtx = NULL;
    AVCodec           *pCodec = NULL;
    AVFrame           *pFrame = NULL;
    AVFrame           *pFrameRGB = NULL;
    AVPacket          packet;
    int               frameFinished;
    int               numBytes;
    uint8_t           *buffer = NULL;
    struct SwsContext *sws_ctx = NULL;

    // Register all formats and codecs
    //av_register_all();
    //debug程序需要将test.flv放在对应的project目录下，跟引用的ffmpeg的dll库同一目录
    char filepath[] = "rtsp://183.59.168.27/PLTV/88888905/224/3221227255/10000100000000060000000001066420_0.smil?icip=88888888";
    // Open video file
    if (avformat_open_input(&pFormatCtx, filepath, NULL, NULL) != 0)
        return -1; // Couldn't open file
    // Retrieve stream information
    if (avformat_find_stream_info(pFormatCtx, NULL) < 0)
        return -1; // Couldn't find stream information

    // Dump information about file onto standard error
    av_dump_format(pFormatCtx, 0, filepath, 0);

    // Find the first video stream
    videoStream = -1;
    for (i = 0; i < pFormatCtx->nb_streams; i++)
        if (pFormatCtx->streams[i]->codecpar->codec_type == AVMEDIA_TYPE_VIDEO) {
            videoStream = i;
            break;
        }
    if (videoStream == -1)
        return -1; // Didn't find a video stream

    // Get a pointer to the codec context for the video stream
    pCodecpar = pFormatCtx->streams[videoStream]->codecpar;
    // Find the decoder for the video stream
    pCodec = avcodec_find_decoder(pFormatCtx->streams[videoStream]->codecpar->codec_id);
    if (pCodec == NULL) {
        fprintf(stderr, "Unsupported codec!\n");
        return -1; // Codec not found
    }

    // Copy context
    pCodecCtx = avcodec_alloc_context3(pCodec);
    if (avcodec_parameters_to_context(pCodecCtx, pCodecpar) != 0) {
        fprintf(stderr, "Couldn't copy codec context");
        return -1; // Error copying codec context
    }

    // Open codec
    if (avcodec_open2(pCodecCtx, pCodec, NULL) < 0)
        return -1; // Could not open codec

    // Allocate video frame
    pFrame = av_frame_alloc();

    // Allocate an AVFrame structure
    pFrameRGB = av_frame_alloc();
    if (pFrameRGB == NULL)
        return -1;

    // Determine required buffer size and allocate buffer
    numBytes = av_image_get_buffer_size(AV_PIX_FMT_RGB24, pCodecCtx->width,
        pCodecCtx->height, 1);
    buffer = (uint8_t *)av_malloc(numBytes*sizeof(uint8_t));

    // Assign appropriate parts of buffer to image planes in pFrameRGB
    // Note that pFrameRGB is an AVFrame, but AVFrame is a superset
    // of AVPicture
    av_image_fill_arrays(pFrameRGB->data, pFrameRGB->linesize, buffer, AV_PIX_FMT_RGB24,
        pCodecCtx->width, pCodecCtx->height, 1);

    // initialize SWS context for software scaling
    sws_ctx = sws_getContext(pCodecCtx->width,
        pCodecCtx->height,
        pCodecCtx->pix_fmt,
        pCodecCtx->width,
        pCodecCtx->height,
        AV_PIX_FMT_RGB24,
        SWS_BILINEAR,
        NULL,
        NULL,
        NULL
        );

    // Read frames and save first five frames to disk
    i = 0;
    while (av_read_frame(pFormatCtx, &packet) >= 0) {
        // Is this a packet from the video stream?
        if (packet.stream_index == videoStream) {
            // Decode video frame
            avcodec_send_packet(pCodecCtx, &packet);
            if (avcodec_receive_frame(pCodecCtx, pFrame) != 0)
                continue;

            // Convert the image from its native format to RGB
            sws_scale(sws_ctx, (uint8_t const * const *)pFrame->data,
                pFrame->linesize, 0, pCodecCtx->height,
                pFrameRGB->data, pFrameRGB->linesize);

            // Save the frame to disk
            if (++i <= 10)
                SaveFrame(pFrameRGB, pCodecCtx->width, pCodecCtx->height, i);
        }

        // Free the packet that was allocated by av_read_frame
        av_packet_unref(&packet);
    }

    // Free the RGB image
    av_free(buffer);
    av_frame_free(&pFrameRGB);

    // Free the YUV frame
    av_frame_free(&pFrame);

    // Close the codecs
    avcodec_close(pCodecCtx);

    // Close the video file
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

//package main
//
//import (
//	"log"
//
//	"github.com/giorgisio/goav/avcodec"
//	"github.com/giorgisio/goav/avdevice"
//	"github.com/giorgisio/goav/avfilter"
//	"github.com/giorgisio/goav/avformat"
//	"github.com/giorgisio/goav/avutil"
//	"github.com/giorgisio/goav/swresample"
//	"github.com/giorgisio/goav/swscale"
//)
//
//func main() {
//
//	// Register all formats and codecs
//	avformat.AvRegisterAll()
//	avcodec.AvcodecRegisterAll()
//
//	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
//	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
//	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
//	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
//	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
//	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())
//
//}
