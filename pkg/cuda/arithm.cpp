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
#include "arithm.h"
#include <string.h>

void GpuAbs(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::abs(*src, *dst);
        return;
    }
    cv::cuda::abs(*src, *dst, *s);
}

void GpuAbsDiff(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::absdiff(*src1, *src2, *dst);
        return;
    }
    cv::cuda::absdiff(*src1, *src2, *dst, *s);
}

void GpuAdd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::add(*src1, *src2, *dst);
        return;
    }
    cv::cuda::add(*src1, *src2, *dst, cv::noArray(), -1, *s);
}

void GpuBitwiseAnd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::bitwise_and(*src1, *src2, *dst);
        return;
    }
    cv::cuda::bitwise_and(*src1, *src2, *dst, cv::noArray(), *s);
}

void GpuBitwiseNot(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::bitwise_not(*src, *dst);
        return;
    }
    cv::cuda::bitwise_not(*src, *dst, cv::noArray(), *s);
}

void GpuBitwiseOr(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::bitwise_or(*src1, *src2, *dst);
        return;
    }
    cv::cuda::bitwise_or(*src1, *src2, *dst, cv::noArray(), *s);
}

void GpuBitwiseXor(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::bitwise_xor(*src1, *src2, *dst);
        return;
    }
    cv::cuda::bitwise_xor(*src1, *src2, *dst, cv::noArray(), *s);
}

void GpuDivide(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::divide(*src1, *src2, *dst);
        return;
    }
    cv::cuda::divide(*src1, *src2, *dst, 1, -1, *s);
}

void GpuExp(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::exp(*src, *dst);
        return;
    }
    cv::cuda::exp(*src, *dst, *s);
}

void GpuLog(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::log(*src, *dst);
        return;
    }
    cv::cuda::log(*src, *dst, *s);
}

void GpuMax(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::max(*src1, *src2, *dst);
        return;
    }
    cv::cuda::max(*src1, *src2, *dst, *s);
}

void GpuMin(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::min(*src1, *src2, *dst);
        return;
    }
    cv::cuda::min(*src1, *src2, *dst, *s);
}

void GpuMultiply(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::multiply(*src1, *src2, *dst);
        return;
    }
    cv::cuda::multiply(*src1, *src2, *dst, 1, -1, *s);
}

void GpuSqr(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::sqr(*src, *dst);
        return;
    }
    cv::cuda::sqr(*src, *dst, *s);
}

void GpuSqrt(GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::sqrt(*src, *dst);
        return;
    }
    cv::cuda::sqrt(*src, *dst, *s);
}

void GpuSubtract(GpuMat src1, GpuMat src2, GpuMat dst, Stream s) {
    if (s == NULL) {
        cv::cuda::subtract(*src1, *src2, *dst);
        return;
    }
    cv::cuda::subtract(*src1, *src2, *dst, cv::noArray(), -1, *s);
}

void GpuThreshold(GpuMat src, GpuMat dst, double thresh, double maxval, int typ, Stream s) {
    if (s == NULL) {
        cv::cuda::threshold(*src, *dst, thresh, maxval, typ);
        return;
    }

    cv::cuda::threshold(*src, *dst, thresh, maxval, typ, *s);
}

void GpuFlip(GpuMat src, GpuMat dst, int flipCode, Stream s) {
    if (s == NULL) {
        cv::cuda::flip(*src, *dst, flipCode);
        return;
    }
    cv::cuda::flip(*src, *dst, flipCode, *s);
}