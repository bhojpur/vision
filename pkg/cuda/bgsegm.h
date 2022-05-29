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

#ifndef _OPENCV3_CUDABGSEGM_H_
#define _OPENCV3_CUDABGSEGM_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/cudabgsegm.hpp>

extern "C" {
#endif

#include "../engine/core.h"
#include "cuda.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::cuda::BackgroundSubtractorMOG2>* CudaBackgroundSubtractorMOG2;
typedef cv::Ptr<cv::cuda::BackgroundSubtractorMOG>* CudaBackgroundSubtractorMOG;
#else
typedef void* CudaBackgroundSubtractorMOG2;
typedef void* CudaBackgroundSubtractorMOG;
#endif

CudaBackgroundSubtractorMOG2 CudaBackgroundSubtractorMOG2_Create();
void CudaBackgroundSubtractorMOG2_Close(CudaBackgroundSubtractorMOG2 b);
void CudaBackgroundSubtractorMOG2_Apply(CudaBackgroundSubtractorMOG2 b, GpuMat src, GpuMat dst, Stream s);

CudaBackgroundSubtractorMOG CudaBackgroundSubtractorMOG_Create();
void CudaBackgroundSubtractorMOG_Close(CudaBackgroundSubtractorMOG b);
void CudaBackgroundSubtractorMOG_Apply(CudaBackgroundSubtractorMOG b, GpuMat src, GpuMat dst, Stream s);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDABGSEGM_H_