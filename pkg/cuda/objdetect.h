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

#ifndef _OPENCV3_GPU_H_
#define _OPENCV3_GPU_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/core/cuda.hpp>
#include <opencv2/cudaobjdetect.hpp>

extern "C" {
#endif

#include "../engine/core.h"
#include "cuda.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::cuda::CascadeClassifier>* CascadeClassifier_GPU;
typedef cv::Ptr<cv::cuda::HOG>* HOG;
#else
typedef void* CascadeClassifier_GPU;
typedef void* HOG;
#endif

// CascadeClassifier
CascadeClassifier_GPU CascadeClassifier_GPU_Create(const char*  cascade_name);
struct Rects CascadeClassifier_GPU_DetectMultiScale(CascadeClassifier_GPU cs, GpuMat img);

// HOG
HOG HOG_Create();
HOG HOG_CreateWithParams(Size winSize, Size blockSize, Size blockStride, Size cellSize, int nbins);
struct Rects HOG_DetectMultiScale(HOG hog, GpuMat img);
GpuMat HOG_Compute(HOG hog, GpuMat img);
Mat HOG_GetPeopleDetector(HOG hog);
void HOG_SetSVMDetector(HOG hog, Mat det);
int HOG_GetDescriptorFormat(HOG hog);
size_t HOG_GetBlockHistogramSize(HOG hog);
size_t HOG_GetDescriptorSize(HOG hog);
bool HOG_GetGammaCorrection(HOG hog);
int HOG_GetGroupThreshold(HOG hog);
double HOG_GetHitThreshold(HOG hog);
double HOG_GetL2HysThreshold(HOG hog);
int HOG_GetNumLevels(HOG hog);
double HOG_GetScaleFactor(HOG hog);
double HOG_GetWinSigma(HOG hog);
struct Size HOG_GetWinStride(HOG hog);
void HOG_SetDescriptorFormat(HOG hog, int descrFormat);
void HOG_SetGammaCorrection(HOG hog, bool gammaCorrection);
void HOG_SetGroupThreshold(HOG hog, int groupThreshold);
void HOG_SetHitThreshold(HOG hog, double hitThreshold);
void HOG_SetL2HysThreshold(HOG hog, double thresholdL2hys);
void HOG_SetNumLevels(HOG hog, int nlevels);
void HOG_SetScaleFactor(HOG hog, double scale0);
void HOG_SetWinSigma(HOG hog, double winSigma);
void HOG_SetWinStride(HOG hog, Size dsize);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_GPU_H_