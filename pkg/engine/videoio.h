// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

#ifndef _OPENCV3_VIDEOIO_H_
#define _OPENCV3_VIDEOIO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

#ifdef __cplusplus
typedef cv::VideoCapture* VideoCapture;
typedef cv::VideoWriter* VideoWriter;
#else
typedef void* VideoCapture;
typedef void* VideoWriter;
#endif

// VideoCapture
VideoCapture VideoCapture_New();
void VideoCapture_Close(VideoCapture v);
bool VideoCapture_Open(VideoCapture v, const char* uri);
bool VideoCapture_OpenWithAPI(VideoCapture v, const char* uri, int apiPreference);
bool VideoCapture_OpenDevice(VideoCapture v, int device);
bool VideoCapture_OpenDeviceWithAPI(VideoCapture v, int device, int apiPreference);
void VideoCapture_Set(VideoCapture v, int prop, double param);
double VideoCapture_Get(VideoCapture v, int prop);
int VideoCapture_IsOpened(VideoCapture v);
int VideoCapture_Read(VideoCapture v, Mat buf);
void VideoCapture_Grab(VideoCapture v, int skip);

// VideoWriter
VideoWriter VideoWriter_New();
void VideoWriter_Close(VideoWriter vw);
void VideoWriter_Open(VideoWriter vw, const char* name, const char* codec, double fps, int width,
                      int height, bool isColor);
int VideoWriter_IsOpened(VideoWriter vw);
void VideoWriter_Write(VideoWriter vw, Mat img);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_VIDEOIO_H_