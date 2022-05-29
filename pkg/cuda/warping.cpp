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

#include "warping.h"

void CudaResize(GpuMat src, GpuMat dst, Size dsize, double fx, double fy, int interp, Stream s) {
    cv::Size sz(dsize.width, dsize.height);

    if (s == NULL) {
        cv::cuda::resize(*src, *dst, sz, fx, fy, interp);
        return;
    }
    cv::cuda::resize(*src, *dst, sz, fx, fy, interp, *s);
}

void CudaPyrDown(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::pyrDown(*src, *dst);
        return;
    }
    cv::cuda::pyrDown(*src, *dst, *s);
}

void CudaPyrUp(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::pyrUp(*src, *dst);
        return;
    }
    cv::cuda::pyrUp(*src, *dst, *s);
}

void CudaBuildWarpAffineMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap, Stream s) {
    cv::Size sz(dsize.width, dsize.height);
    if (s == NULL) {
        cv::cuda::buildWarpAffineMaps(*M, inverse, sz, *xmap, *ymap);
        return;
    }
    cv::cuda::buildWarpAffineMaps(*M, inverse, sz, *xmap, *ymap, *s);
}

void CudaBuildWarpPerspectiveMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap, Stream s) {
    cv::Size sz(dsize.width, dsize.height);
    if (s == NULL) {
        cv::cuda::buildWarpPerspectiveMaps(*M, inverse, sz, *xmap, *ymap);
        return;
    }
    cv::cuda::buildWarpPerspectiveMaps(*M, inverse, sz, *xmap, *ymap, *s);
}

void CudaRemap(GpuMat src, GpuMat dst, GpuMat xmap, GpuMat ymap, int interp, int borderMode, Scalar borderValue, Stream s) {
    cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
    if (s == NULL) {
        cv::cuda::remap(*src, *dst, *xmap, *ymap, interp, borderMode, c);
        return;
    }
    cv::cuda::remap(*src, *dst, *xmap, *ymap, interp, borderMode, c, *s);
}

void CudaRotate(GpuMat src, GpuMat dst, Size dsize, double angle, double xShift, double yShift, int interp, Stream s) {
    cv::Size sz(dsize.width, dsize.height);
    if (s == NULL) {
        cv::cuda::rotate(*src, *dst, sz, angle, xShift, yShift, interp);
        return;
    }
    cv::cuda::rotate(*src, *dst, sz, angle, xShift, yShift, interp, *s);
}

void CudaWarpAffine(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue, Stream s) {
    cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
    cv::Size sz(dsize.width, dsize.height);

    if (s == NULL) {
        cv::cuda::warpAffine(*src, *dst, *M, sz, flags, borderMode, c);
        return;
    }
    cv::cuda::warpAffine(*src, *dst, *M, sz, flags, borderMode, c, *s);
}

void CudaWarpPerspective(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue, Stream s) {
    cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
    cv::Size sz(dsize.width, dsize.height);
    if (s == NULL) {
        cv::cuda::warpPerspective(*src, *dst, *M, sz, flags, borderMode, c);
        return;
    }
    cv::cuda::warpPerspective(*src, *dst, *M, sz, flags, borderMode, c, *s);
}