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

import (
	"image"
	"testing"

	engine "github.com/bhojpur/vision/pkg/engine"
)

func TestGaussianFilter_Apply(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in GaussianFilter test")
	}
	defer src.Close()

	cimg, dimg := NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(src)

	filter := NewGaussianFilter(src.Type(), src.Type(), image.Pt(23, 23), 30)
	defer filter.Close()

	filter.Apply(cimg, &dimg)

	dest := engine.NewMat()
	defer dest.Close()

	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty GaussianFilter test")
	}
	if src.Rows() != dest.Rows() {
		t.Error("Invalid GaussianFilter test rows")
	}
	if src.Cols() != dest.Cols() {
		t.Error("Invalid GaussianFilter test cols")
	}
}

func TestGaussianFilter_ApplyWithStream(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in GaussianFilter test")
	}
	defer src.Close()

	cimg, dimg := NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	filter := NewGaussianFilter(src.Type(), src.Type(), image.Pt(23, 23), 30)
	defer filter.Close()

	stream := NewStream()
	defer stream.Close()

	dest := engine.NewMat()
	defer dest.Close()

	cimg.UploadWithStream(src, stream)
	filter.ApplyWithStream(cimg, &dimg, stream)
	dimg.DownloadWithStream(&dest, stream)

	stream.WaitForCompletion()

	if dest.Empty() {
		t.Error("Empty GaussianFilter test")
	}
	if src.Rows() != dest.Rows() {
		t.Error("Invalid GaussianFilter test rows")
	}
	if src.Cols() != dest.Cols() {
		t.Error("Invalid GaussianFilter test cols")
	}
}

func TestSobelFilter_Apply(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in SobelFilter test")
	}
	defer src.Close()

	cimg, dimg := NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	filter := NewSobelFilter(src.Type(), src.Type(), 0, 1)
	defer filter.Close()

	dest := engine.NewMat()
	defer dest.Close()

	cimg.Upload(src)
	filter.Apply(cimg, &dimg)
	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty SobelFilter test")
	}
	if src.Rows() != dest.Rows() {
		t.Error("Invalid SobelFilter test rows")
	}
	if src.Cols() != dest.Cols() {
		t.Error("Invalid SobelFilter test cols")
	}
}

func TestSobelFilter_ApplyWithStream(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in SobelFilter test")
	}
	defer src.Close()

	cimg, dimg := NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	filter := NewSobelFilter(src.Type(), src.Type(), 0, 1)
	defer filter.Close()

	stream := NewStream()
	defer stream.Close()

	dest := engine.NewMat()
	defer dest.Close()

	cimg.UploadWithStream(src, stream)
	filter.ApplyWithStream(cimg, &dimg, stream)
	dimg.DownloadWithStream(&dest, stream)

	stream.WaitForCompletion()

	if dest.Empty() {
		t.Error("Empty SobelFilter test")
	}
	if src.Rows() != dest.Rows() {
		t.Error("Invalid SobelFilter test rows")
	}
	if src.Cols() != dest.Cols() {
		t.Error("Invalid SobelFilter test cols")
	}
}
