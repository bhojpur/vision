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

#ifndef _OPENCV3_CUDA_ARITHM_H_
#define _OPENCV3_CUDA_ARITHM_H_

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/cudaarithm.hpp>
extern "C" {
#endif
#include "cuda.h"

void GpuAbs(GpuMat src, GpuMat dst, Stream s);
void GpuAbsDiff(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuAdd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuBitwiseAnd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuBitwiseNot(GpuMat src, GpuMat dst, Stream s);
void GpuBitwiseOr(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuBitwiseXor(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuDivide(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuExp(GpuMat src, GpuMat dst, Stream s);
void GpuLog(GpuMat src, GpuMat dst, Stream s);
void GpuMax(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuMin(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuMultiply(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuSqr(GpuMat src, GpuMat dst, Stream s);
void GpuSqrt(GpuMat src, GpuMat dst, Stream s);
void GpuSubtract(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuThreshold(GpuMat src, GpuMat dst, double thresh, double maxval, int typ, Stream s);
void GpuFlip(GpuMat src, GpuMat dst, int flipCode, Stream s);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_ARITHM_H_