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
	"image/color"
	"math"
	"testing"

	engine "github.com/bhojpur/vision/pkg/engine"
)

func TestResize(t *testing.T) {
	src := engine.IMRead("../../images/logo.png", engine.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Resize test")
	}
	defer src.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(src)

	dst := engine.NewMat()
	defer dst.Close()

	Resize(cimg, &dimg, image.Point{}, 0.5, 0.5, InterpolationDefault)
	dimg.Download(&dst)
	if dst.Cols() != 200 || dst.Rows() != 172 {
		t.Errorf("Expected dst size of 200x172 got %dx%d", dst.Cols(), dst.Rows())
	}

	Resize(cimg, &dimg, image.Pt(440, 377), 0, 0, InterpolationCubic)
	dimg.Download(&dst)
	if dst.Cols() != 440 || dst.Rows() != 377 {
		t.Errorf("Expected dst size of 440x377 got %dx%d", dst.Cols(), dst.Rows())
	}
}

func TestResizeWithStream(t *testing.T) {
	src := engine.IMRead("../../images/logo.png", engine.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Resize test")
	}
	defer src.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(src)

	dst := engine.NewMat()
	defer dst.Close()

	stream := NewStream()
	defer stream.Close()

	ResizeWithStream(cimg, &dimg, image.Point{}, 0.5, 0.5, InterpolationDefault, stream)
	dimg.Download(&dst)
	if dst.Cols() != 200 || dst.Rows() != 172 {
		t.Errorf("Expected dst size of 200x172 got %dx%d", dst.Cols(), dst.Rows())
	}

	ResizeWithStream(cimg, &dimg, image.Pt(440, 377), 0, 0, InterpolationCubic, stream)
	dimg.Download(&dst)
	if dst.Cols() != 440 || dst.Rows() != 377 {
		t.Errorf("Expected dst size of 440x377 got %dx%d", dst.Cols(), dst.Rows())
	}
}

func TestPyrDown(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in PyrDown test")
	}
	defer src.Close()

	var gsrc, gdst = NewGpuMat(), NewGpuMat()
	defer gsrc.Close()
	defer gdst.Close()

	gsrc.Upload(src)

	dst := engine.NewMat()
	defer dst.Close()

	PyrDown(gsrc, &gdst)
	gdst.Download(&dst)
	if dst.Empty() && math.Abs(float64(src.Cols()-2*dst.Cols())) < 2.0 && math.Abs(float64(src.Rows()-2*dst.Rows())) < 2.0 {
		t.Error("Invalid PyrDown test")
	}
}

func TestPyrDownWithStream(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in PyrDown test")
	}
	defer src.Close()

	var gsrc, gdst = NewGpuMat(), NewGpuMat()
	defer gsrc.Close()
	defer gdst.Close()

	gsrc.Upload(src)

	dst := engine.NewMat()
	defer dst.Close()

	stream := NewStream()
	defer stream.Close()

	PyrDownWithStream(gsrc, &gdst, stream)
	gdst.Download(&dst)
	if dst.Empty() && math.Abs(float64(src.Cols()-2*dst.Cols())) < 2.0 && math.Abs(float64(src.Rows()-2*dst.Rows())) < 2.0 {
		t.Error("Invalid PyrDown test")
	}
}

func TestPyrUp(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in PyrUp test")
	}
	defer src.Close()

	var gsrc, gdst = NewGpuMat(), NewGpuMat()
	defer gsrc.Close()
	defer gdst.Close()

	gsrc.Upload(src)

	dst := engine.NewMat()
	defer dst.Close()

	PyrDown(gsrc, &gdst)
	if dst.Empty() && math.Abs(float64(2*src.Cols()-dst.Cols())) < 2.0 && math.Abs(float64(2*src.Rows()-dst.Rows())) < 2.0 {
		t.Error("Invalid PyrUp test")
	}
}

func TestPyrUpWithStream(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in PyrUp test")
	}
	defer src.Close()

	var gsrc, gdst = NewGpuMat(), NewGpuMat()
	defer gsrc.Close()
	defer gdst.Close()

	gsrc.Upload(src)

	dst := engine.NewMat()
	defer dst.Close()

	stream := NewStream()
	defer stream.Close()

	PyrDownWithStream(gsrc, &gdst, stream)
	if dst.Empty() && math.Abs(float64(2*src.Cols()-dst.Cols())) < 2.0 && math.Abs(float64(2*src.Rows()-dst.Rows())) < 2.0 {
		t.Error("Invalid PyrUp test")
	}
}

func TestRemap(t *testing.T) {
	src := engine.IMRead("../../images/logo.png", engine.IMReadUnchanged)
	defer src.Close()

	dst := engine.NewMat()
	defer dst.Close()

	map1 := engine.NewMatWithSize(256, 256, engine.MatTypeCV32F)
	defer map1.Close()
	map1.SetFloatAt(50, 50, 25.4)
	map2 := engine.NewMatWithSize(256, 256, engine.MatTypeCV32F)
	defer map2.Close()

	gsrc, gdst, gmap1, gmap2 := NewGpuMat(), NewGpuMat(), NewGpuMat(), NewGpuMat()
	gsrc.Upload(src)
	gmap1.Upload(map1)
	gmap2.Upload(map2)
	Remap(gsrc, &gdst, &gmap1, &gmap2, InterpolationDefault, BorderConstant, color.RGBA{0, 0, 0, 0})
	gdst.Download(&dst)
	if ok := dst.Empty(); ok {
		t.Errorf("Remap(): dst is empty")
	}
}

func TestRemapWithStream(t *testing.T) {
	src := engine.IMRead("../../images/logo.png", engine.IMReadUnchanged)
	defer src.Close()

	dst := engine.NewMat()
	defer dst.Close()

	map1 := engine.NewMatWithSize(256, 256, engine.MatTypeCV32F)
	defer map1.Close()
	map1.SetFloatAt(50, 50, 25.4)
	map2 := engine.NewMatWithSize(256, 256, engine.MatTypeCV32F)
	defer map2.Close()

	gsrc, gdst, gmap1, gmap2 := NewGpuMat(), NewGpuMat(), NewGpuMat(), NewGpuMat()
	gsrc.Upload(src)
	gmap1.Upload(map1)
	gmap2.Upload(map2)

	stream := NewStream()
	defer stream.Close()

	RemapWithStream(gsrc, &gdst, &gmap1, &gmap2, InterpolationDefault, BorderConstant, color.RGBA{0, 0, 0, 0}, stream)
	gdst.Download(&dst)
	if ok := dst.Empty(); ok {
		t.Errorf("Remap(): dst is empty")
	}
}
