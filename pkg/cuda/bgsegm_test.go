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
	"testing"

	engine "github.com/bhojpur/vision/pkg/engine"
)

func TestCudaMOG2(t *testing.T) {
	img := engine.IMRead("../../images/face.jpg", engine.IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in CudaMOG2 test")
	}
	defer img.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(img)

	dst := engine.NewMat()
	defer dst.Close()

	mog2 := NewBackgroundSubtractorMOG2()
	defer mog2.Close()

	mog2.Apply(cimg, &dimg)

	dimg.Download(&dst)

	if dst.Empty() {
		t.Error("Error in TestCudaMOG2 test")
	}
}

func TestCudaMOG2WithStream(t *testing.T) {
	img := engine.IMRead("../../images/face.jpg", engine.IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in CudaMOG2 test")
	}
	defer img.Close()

	var cimg, dimg, s = NewGpuMat(), NewGpuMat(), NewStream()
	defer cimg.Close()
	defer dimg.Close()
	defer s.Close()

	dst := engine.NewMat()
	defer dst.Close()

	mog2 := NewBackgroundSubtractorMOG2()
	defer mog2.Close()

	cimg.UploadWithStream(img, s)
	mog2.ApplyWithStream(cimg, &dimg, s)
	dimg.DownloadWithStream(&dst, s)

	s.WaitForCompletion()

	if dst.Empty() {
		t.Error("Error in TestCudaMOG2 test")
	}
}

func TestCudaMOG(t *testing.T) {
	img := engine.IMRead("../../images/face.jpg", engine.IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in CudaMOG test")
	}
	defer img.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(img)

	dst := engine.NewMat()
	defer dst.Close()

	mog2 := NewBackgroundSubtractorMOG()
	defer mog2.Close()

	mog2.Apply(cimg, &dimg)

	dimg.Download(&dst)

	if dst.Empty() {
		t.Error("Error in TestCudaMOG test")
	}
}

func TestCudaMOGWithStream(t *testing.T) {
	img := engine.IMRead("../../images/face.jpg", engine.IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in CudaMOG test")
	}
	defer img.Close()

	var cimg, dimg, s = NewGpuMat(), NewGpuMat(), NewStream()
	defer cimg.Close()
	defer dimg.Close()
	defer s.Close()

	dst := engine.NewMat()
	defer dst.Close()

	mog2 := NewBackgroundSubtractorMOG()
	defer mog2.Close()

	cimg.UploadWithStream(img, s)
	mog2.ApplyWithStream(cimg, &dimg, s)
	dimg.DownloadWithStream(&dst, s)

	s.WaitForCompletion()

	if dst.Empty() {
		t.Error("Error in TestCudaMOG test")
	}
}
