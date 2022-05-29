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

#ifndef _OPENCV3_CUDAWARPING_H_
#define _OPENCV3_CUDAWARPING_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/cudawarping.hpp>

extern "C" {
#endif

#include "../engine/core.h"
#include "cuda.h"

void CudaResize(GpuMat src, GpuMat dst, Size dsize, double fx, double fy, int interp, Stream s);
void CudaPyrDown(GpuMat src, GpuMat dst, Stream s);
void CudaPyrUp(GpuMat src, GpuMat dst, Stream s);
void CudaBuildWarpAffineMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap, Stream s);
void CudaBuildWarpPerspectiveMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap, Stream s);
void CudaRemap(GpuMat src, GpuMat dst, GpuMat xmap, GpuMat ymap, int interp, int borderMode, Scalar borderValue, Stream s);
void CudaRotate(GpuMat src, GpuMat dst, Size dsize, double angle, double xShift, double yShift, int interp, Stream s);
void CudaWarpAffine(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue, Stream s);
void CudaWarpPerspective(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue, Stream s);
#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDAWARPING_H_