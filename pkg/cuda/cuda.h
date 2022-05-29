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

#ifndef _OPENCV3_CUDA_H_
#define _OPENCV3_CUDA_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/core/cuda.hpp>

extern "C" {
#endif

#include "../engine/core.h"

#ifdef __cplusplus
typedef cv::cuda::GpuMat* GpuMat;
typedef cv::cuda::Stream* Stream;
#else
typedef void* GpuMat;
typedef void* Stream;
#endif

GpuMat GpuMat_New();
GpuMat GpuMat_NewFromMat(Mat mat);
GpuMat GpuMat_NewWithSize(int rows, int cols, int type);
void GpuMat_Upload(GpuMat m, Mat data, Stream s);
void GpuMat_Download(GpuMat m, Mat dst, Stream s);
void GpuMat_Close(GpuMat m);
int GpuMat_Empty(GpuMat m);
void GpuMat_ConvertTo(GpuMat m, GpuMat dst, int type, Stream s);
void GpuMat_CopyTo(GpuMat m, GpuMat dst, Stream s);
GpuMat GpuMat_Reshape(GpuMat m, int cn, int rows);
int GpuMat_Cols(GpuMat m);
int GpuMat_Rows(GpuMat m);
int GpuMat_Channels(GpuMat m);
int GpuMat_Type(GpuMat m);

void PrintCudaDeviceInfo(int device);
void PrintShortCudaDeviceInfo(int device);
int GetCudaEnabledDeviceCount();
int GetCudaDevice();
void SetCudaDevice(int device);
void ResetCudaDevice();

Stream Stream_New();
void Stream_Close(Stream s);
bool Stream_QueryIfComplete(Stream s);
void Stream_WaitForCompletion(Stream s);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_H_