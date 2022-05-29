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
#include "../engine/core.h"
#include "core.h"
#include "filters.h"
*/
import "C"
import (
	"image"
	"unsafe"

	engine "github.com/bhojpur/vision/pkg/engine"
)

// GaussianFilter
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d66/group__cudafilters.html#gaa4df286369114cfd4b144ae211f6a6c8
//
type GaussianFilter struct {
	p unsafe.Pointer
}

// NewGaussianFilter returns a new GaussianFilter.
func NewGaussianFilter(srcType engine.MatType, dstType engine.MatType, ksize image.Point, sigma1 float64) GaussianFilter {
	pSize := C.struct_Size{
		width:  C.int(ksize.X),
		height: C.int(ksize.Y),
	}

	return GaussianFilter{p: unsafe.Pointer(C.CreateGaussianFilter(C.int(srcType), C.int(dstType), pSize, C.double(sigma1)))}
}

// Close GaussianFilter
func (gf *GaussianFilter) Close() error {
	C.GaussianFilter_Close((C.GaussianFilter)(gf.p))
	gf.p = nil
	return nil
}

// Apply applies the Gaussian filter.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d2b/classcv_1_1cuda_1_1Filter.html#a20b58d13871027473b4c39cc698cf80f
//
func (gf *GaussianFilter) Apply(img GpuMat, dst *GpuMat) {
	C.GaussianFilter_Apply(C.GaussianFilter(gf.p), img.p, dst.p, nil)
	return
}

// ApplyWithStream applies the Gaussian filter
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d2b/classcv_1_1cuda_1_1Filter.html#a20b58d13871027473b4c39cc698cf80f
//
func (gf *GaussianFilter) ApplyWithStream(img GpuMat, dst *GpuMat, s Stream) {
	C.GaussianFilter_Apply(C.GaussianFilter(gf.p), img.p, dst.p, s.p)
	return
}

// SobelFilter
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d66/group__cudafilters.html#gabf85fe61958bb21e93211a6fcc7c5c3b
//
type SobelFilter struct {
	p unsafe.Pointer
}

// NewSobelFilter returns a new SobelFilter.
func NewSobelFilter(srcType engine.MatType, dstType engine.MatType, dx int, dy int) SobelFilter {
	return SobelFilter{p: unsafe.Pointer(C.CreateSobelFilter(C.int(srcType), C.int(dstType), C.int(dx), C.int(dy)))}
}

// NewSobelFilterWithParams returns a new SobelFilter.
func NewSobelFilterWithParams(srcType engine.MatType, dstType engine.MatType, dx int, dy int, ksize int, scale float64, rowBorderMode int, columnBorderMode int) SobelFilter {
	return SobelFilter{p: unsafe.Pointer(C.CreateSobelFilterWithParams(C.int(srcType), C.int(dstType), C.int(dx), C.int(dy), C.int(ksize), C.double(scale), C.int(rowBorderMode), C.int(columnBorderMode)))}
}

// Close SobelFilter
func (sf *SobelFilter) Close() error {
	C.SobelFilter_Close((C.SobelFilter)(sf.p))
	sf.p = nil
	return nil
}

// Apply applies the Sobel filter.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d2b/classcv_1_1cuda_1_1Filter.html#a20b58d13871027473b4c39cc698cf80f
//
func (sf *SobelFilter) Apply(img GpuMat, dst *GpuMat) {
	C.SobelFilter_Apply(C.SobelFilter(sf.p), img.p, dst.p, nil)
	return
}

// ApplyWithStream applies the Sobel filter
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d2b/classcv_1_1cuda_1_1Filter.html#a20b58d13871027473b4c39cc698cf80f
//
func (sf *SobelFilter) ApplyWithStream(img GpuMat, dst *GpuMat, s Stream) {
	C.SobelFilter_Apply(C.SobelFilter(sf.p), img.p, dst.p, s.p)
	return
}
