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

#include "bgsegm.h"

CudaBackgroundSubtractorMOG2 CudaBackgroundSubtractorMOG2_Create() {
    return new cv::Ptr<cv::cuda::BackgroundSubtractorMOG2>(cv::cuda::createBackgroundSubtractorMOG2());
}

void CudaBackgroundSubtractorMOG2_Close(CudaBackgroundSubtractorMOG2 b) {
    delete b;
}

void CudaBackgroundSubtractorMOG2_Apply(CudaBackgroundSubtractorMOG2 b, GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        (*b)->apply(*src, *dst);
        return;
    }
    (*b)->apply(*src, *dst, -1.0, *s);
}

CudaBackgroundSubtractorMOG CudaBackgroundSubtractorMOG_Create() {
    return new cv::Ptr<cv::cuda::BackgroundSubtractorMOG>(cv::cuda::createBackgroundSubtractorMOG());
}

void CudaBackgroundSubtractorMOG_Close(CudaBackgroundSubtractorMOG b) {
    delete b;
}

void CudaBackgroundSubtractorMOG_Apply(CudaBackgroundSubtractorMOG b, GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        (*b)->apply(*src, *dst);
        return;
    }
    (*b)->apply(*src, *dst, -1.0, *s);
}