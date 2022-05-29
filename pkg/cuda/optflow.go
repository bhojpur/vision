package cuda

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

/*
#include <stdlib.h>
#include "cuda.h"
#include "optflow.h"
*/
import "C"
import "unsafe"

// SparsePyrLKOpticalFlow is a wrapper around the cv::cuda::SparsePyrLKOpticalFlow.
type SparsePyrLKOpticalFlow struct {
	// C.SparsePyrLKOpticalFlow
	p unsafe.Pointer
}

// NewSparsePyrLKOpticalFlow returns a new SparsePyrLKOpticalFlow
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d05/classcv_1_1cuda_1_1SparsePyrLKOpticalFlow.html#a6bcd2d457532d7db76c3e7f11b60063b
//
func NewSparsePyrLKOpticalFlow() SparsePyrLKOpticalFlow {
	return SparsePyrLKOpticalFlow{p: unsafe.Pointer(C.CudaSparsePyrLKOpticalFlow_Create())}
}

// Calc calculates a sparse optical flow.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/dcf/classcv_1_1cuda_1_1SparseOpticalFlow.html#a80d5efbb7788e3dc4c49e6226ba34347
func (s SparsePyrLKOpticalFlow) Calc(prevImg, nextImg, prevPts, nextPts, status GpuMat) {
	C.CudaSparsePyrLKOpticalFlow_Calc(C.CudaSparsePyrLKOpticalFlow(s.p), prevImg.p, nextImg.p, prevPts.p, nextPts.p, status.p)
}
