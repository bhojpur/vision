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
#include "bgsegm.h"
*/
import "C"
import "unsafe"

// BackgroundSubtractorMOG2 is a wrapper around the cv::cuda::BackgroundSubtractorMOG2.
type BackgroundSubtractorMOG2 struct {
	// C.BackgroundSubtractorMOG2
	p unsafe.Pointer
}

// BackgroundSubtractorMOG is a wrapper around the cv::cuda::BackgroundSubtractorMOG.
type BackgroundSubtractorMOG struct {
	// C.BackgroundSubtractorMOG
	p unsafe.Pointer
}

// NewBackgroundSubtractorMOG2 returns a new BackgroundSubtractor algorithm
// of type MOG2. MOG2 is a Gaussian Mixture-based Background/Foreground
// Segmentation Algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d3d/cudabgsegm_8hpp.html
//
func NewBackgroundSubtractorMOG2() BackgroundSubtractorMOG2 {
	return BackgroundSubtractorMOG2{p: unsafe.Pointer(C.CudaBackgroundSubtractorMOG2_Create())}
}

// Close BackgroundSubtractorMOG2.
func (b *BackgroundSubtractorMOG2) Close() error {
	C.CudaBackgroundSubtractorMOG2_Close((C.CudaBackgroundSubtractorMOG2)(b.p))
	b.p = nil
	return nil
}

// Apply computes a foreground mask using the current BackgroundSubtractorMOG2.
//
// For further details, please see:
// https://docs.opencv.org/master/df/d23/classcv_1_1cuda_1_1BackgroundSubtractorMOG2.html#a92408f07bf1268c1b778cb186b3113b0
//
func (b *BackgroundSubtractorMOG2) Apply(src GpuMat, dst *GpuMat) {
	C.CudaBackgroundSubtractorMOG2_Apply((C.CudaBackgroundSubtractorMOG2)(b.p), src.p, dst.p, nil)
	return
}

// ApplyWithStream computes a foreground mask using the current BackgroundSubtractorMOG2
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/df/d23/classcv_1_1cuda_1_1BackgroundSubtractorMOG2.html#a92408f07bf1268c1b778cb186b3113b0
//
func (b *BackgroundSubtractorMOG2) ApplyWithStream(src GpuMat, dst *GpuMat, s Stream) {
	C.CudaBackgroundSubtractorMOG2_Apply((C.CudaBackgroundSubtractorMOG2)(b.p), src.p, dst.p, s.p)
	return
}

// NewBackgroundSubtractorMOG returns a new BackgroundSubtractor algorithm
// of type MOG. MOG is a Gaussian Mixture-based Background/Foreground
// Segmentation Algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d3d/cudabgsegm_8hpp.html
//
func NewBackgroundSubtractorMOG() BackgroundSubtractorMOG {
	return BackgroundSubtractorMOG{p: unsafe.Pointer(C.CudaBackgroundSubtractorMOG_Create())}
}

// Close BackgroundSubtractorMOG.
func (b *BackgroundSubtractorMOG) Close() error {
	C.CudaBackgroundSubtractorMOG_Close((C.CudaBackgroundSubtractorMOG)(b.p))
	b.p = nil
	return nil
}

// Apply computes a foreground mask using the current BackgroundSubtractorMOG.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/dfe/classcv_1_1cuda_1_1BackgroundSubtractorMOG.html#a8f52d2f7abd1c77c84243efc53972cbf
//
func (b *BackgroundSubtractorMOG) Apply(src GpuMat, dst *GpuMat) {
	C.CudaBackgroundSubtractorMOG_Apply((C.CudaBackgroundSubtractorMOG)(b.p), src.p, dst.p, nil)
	return
}

// ApplyWithStream computes a foreground mask using the current BackgroundSubtractorMOG
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/dfe/classcv_1_1cuda_1_1BackgroundSubtractorMOG.html#a8f52d2f7abd1c77c84243efc53972cbf
//
func (b *BackgroundSubtractorMOG) ApplyWithStream(src GpuMat, dst *GpuMat, s Stream) {
	C.CudaBackgroundSubtractorMOG_Apply((C.CudaBackgroundSubtractorMOG)(b.p), src.p, dst.p, s.p)
	return
}
