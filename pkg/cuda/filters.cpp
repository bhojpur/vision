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

#include "../engine/core.h"
#include "filters.h"
#include <string.h>

GaussianFilter CreateGaussianFilter(int srcType, int dstType, Size ksize, double sigma1) {
    cv::Size sz(ksize.width, ksize.height);
    return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createGaussianFilter(srcType, dstType, sz, sigma1));
}

GaussianFilter CreateGaussianFilterWithParams(int srcType, int dstType, Size ksize, double sigma1, double sigma2, int rowBorderMode, int columnBorderMode) {
    cv::Size sz(ksize.width, ksize.height);
    return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createGaussianFilter(srcType, dstType, sz, sigma1, sigma2, rowBorderMode, columnBorderMode));
}

void GaussianFilter_Close(GaussianFilter gf) {
    delete gf;
}

void GaussianFilter_Apply(GaussianFilter gf, GpuMat img, GpuMat dst, Stream s) {
    if (s == NULL) {
        (*gf)->apply(*img, *dst);
    } else {
        (*gf)->apply(*img, *dst, *s);
    }
    return;
}

SobelFilter CreateSobelFilter(int srcType, int dstType, int dx, int dy) {
    return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createSobelFilter(srcType, dstType, dx, dy));
}

SobelFilter CreateSobelFilterWithParams(int srcType, int dstType, int dx, int dy, int ksize, double scale, int rowBorderMode, int columnBorderMode) {
    return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createSobelFilter(srcType, dstType, dx, dy, ksize, rowBorderMode, columnBorderMode));
}

void SobelFilter_Close(SobelFilter sf) {
    delete sf;
}

void SobelFilter_Apply(SobelFilter sf, GpuMat img, GpuMat dst, Stream s) {
    if (s == NULL) {
        (*sf)->apply(*img, *dst);
    } else {
        (*sf)->apply(*img, *dst, *s);
    }

    return;
}