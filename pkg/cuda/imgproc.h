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

#ifndef _OPENCV3_CUDA_IMGPROC_H_
#define _OPENCV3_CUDA_IMGPROC_H_

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/cudaimgproc.hpp>
#include <opencv2/cudaarithm.hpp>
extern "C" {
#endif
#include "cuda.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::cuda::CannyEdgeDetector>* CannyEdgeDetector;
typedef cv::Ptr<cv::cuda::HoughLinesDetector>* HoughLinesDetector;
typedef cv::Ptr<cv::cuda::HoughSegmentDetector>* HoughSegmentDetector;
#else
typedef void* CannyEdgeDetector;
typedef void* HoughLinesDetector;
typedef void* HoughSegmentDetector;
#endif

// standalone functions
void GpuCvtColor(GpuMat src, GpuMat dst, int code, Stream s);

// CannyEdgeDetector
CannyEdgeDetector CreateCannyEdgeDetector(double lowThresh, double highThresh);
CannyEdgeDetector CreateCannyEdgeDetectorWithParams(double lowThresh, double highThresh, int appertureSize, bool L2gradient);
void CannyEdgeDetector_Close(CannyEdgeDetector det);
void CannyEdgeDetector_Detect(CannyEdgeDetector det, GpuMat img, GpuMat dst, Stream s);
int CannyEdgeDetector_GetAppertureSize(CannyEdgeDetector det);
double CannyEdgeDetector_GetHighThreshold(CannyEdgeDetector det);
bool CannyEdgeDetector_GetL2Gradient(CannyEdgeDetector det);
double CannyEdgeDetector_GetLowThreshold(CannyEdgeDetector det);
void CannyEdgeDetector_SetAppertureSize(CannyEdgeDetector det, int appertureSize);
void CannyEdgeDetector_SetHighThreshold(CannyEdgeDetector det, double highThresh);
void CannyEdgeDetector_SetL2Gradient(CannyEdgeDetector det, bool L2gradient);
void CannyEdgeDetector_SetLowThreshold(CannyEdgeDetector det, double lowThresh);

// HoughLinesDetector
HoughLinesDetector HoughLinesDetector_Create(double rho, double theta, int threshold);
HoughLinesDetector HoughLinesDetector_CreateWithParams(double rho, double theta, int threshold, bool sort, int maxlines);
void HoughLinesDetector_Close(HoughLinesDetector hld);
void HoughLinesDetector_Detect(HoughLinesDetector hld, GpuMat img, GpuMat dst, Stream s);

// HoughSegmentDetector
HoughSegmentDetector HoughSegmentDetector_Create(double rho, double theta, int minLineLength, int maxLineGap);
void HoughSegmentDetector_Close(HoughSegmentDetector hsd);
void HoughSegmentDetector_Detect(HoughSegmentDetector hsd, GpuMat img, GpuMat dst, Stream s);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_IMGPROC_H_