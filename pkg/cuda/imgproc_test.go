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
	"fmt"
	"math"
	"strconv"
	"testing"

	engine "github.com/bhojpur/vision/pkg/engine"
	"github.com/pascaldekloe/goe/verify"
)

func TestCanny_Detect(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in Canny test")
	}
	defer src.Close()

	cimg, dimg := NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	dest := engine.NewMat()
	defer dest.Close()

	detector := NewCannyEdgeDetector(50, 100)
	defer detector.Close()

	cimg.Upload(src)
	detector.Detect(cimg, &dimg)
	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty Canny test")
	}
	if src.Rows() != dest.Rows() {
		t.Error("Invalid Canny test rows")
	}
	if src.Cols() != dest.Cols() {
		t.Error("Invalid Canny test cols")
	}
}

func TestHoughLines_Calc(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in HoughLines test")
	}
	defer src.Close()

	cimg, mimg, dimg := NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer mimg.Close()
	defer dimg.Close()

	canny := NewCannyEdgeDetector(100, 200)
	defer canny.Close()

	detector := NewHoughLinesDetectorWithParams(1, math.Pi/180, 50, true, 4096)
	defer detector.Close()

	dest := engine.NewMat()
	defer dest.Close()

	cimg.Upload(src)
	canny.Detect(cimg, &mimg)
	detector.Detect(mimg, &dimg)
	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty HoughLines test")
	}

	if dest.Rows() != 2 {
		t.Errorf("Invalid HoughLines test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1588 {
		t.Errorf("Invalid HoughLines test cols: %v", dest.Cols())
	}

	expected := map[float32]float32{
		21:  1.5707964,
		337: 0.034906585,
		85:  1.5707964,
		276: 0,
		329: 0.034906585,
	}

	actual := make(map[float32]float32)
	for i := 0; i < dest.Cols(); i += 2 {
		actual[dest.GetFloatAt(0, i)] = dest.GetFloatAt(0, i+1)
	}

	for k, v := range expected {
		s32 := strconv.FormatFloat(float64(k), 'f', -1, 32)
		verify.Values(t, s32, actual[k], v)
	}
}

func TestHoughLines_CalcWithStream(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in HoughLines test")
	}
	defer src.Close()

	cimg, mimg, dimg := NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer mimg.Close()
	defer dimg.Close()

	stream := NewStream()
	defer stream.Close()

	canny := NewCannyEdgeDetector(100, 200)
	defer canny.Close()

	detector := NewHoughLinesDetectorWithParams(1, math.Pi/180, 50, true, 4096)
	defer detector.Close()

	dest := engine.NewMat()
	defer dest.Close()

	cimg.UploadWithStream(src, stream)
	canny.DetectWithStream(cimg, &mimg, stream)
	detector.DetectWithStream(mimg, &dimg, stream)
	dimg.DownloadWithStream(&dest, stream)

	stream.WaitForCompletion()

	if dest.Empty() {
		t.Error("Empty HoughLines test")
	}

	if dest.Rows() != 2 {
		t.Errorf("Invalid HoughLines test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1588 {
		t.Errorf("Invalid HoughLines test cols: %v", dest.Cols())
	}

	expected := map[float32]float32{
		21:  1.5707964,
		337: 0.034906585,
		85:  1.5707964,
		276: 0,
		329: 0.034906585,
	}

	actual := make(map[float32]float32)
	for i := 0; i < dest.Cols(); i += 2 {
		actual[dest.GetFloatAt(0, i)] = dest.GetFloatAt(0, i+1)
	}

	for k, v := range expected {
		s32 := strconv.FormatFloat(float64(k), 'f', -1, 32)
		verify.Values(t, s32, actual[k], v)
	}
}

func TestHoughSegment_Calc(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in HoughSegment test")
	}
	defer src.Close()

	cimg, mimg, dimg := NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer mimg.Close()
	defer dimg.Close()

	canny := NewCannyEdgeDetector(50, 100)
	defer canny.Close()

	detector := NewHoughSegmentDetector(1, math.Pi/180, 150, 50)
	defer detector.Close()

	dest := engine.NewMat()
	defer dest.Close()

	cimg.Upload(src)
	canny.Detect(cimg, &mimg)
	detector.Detect(mimg, &dimg)
	fimg := dimg.Reshape(0, dimg.Cols())
	defer fimg.Close()
	fimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty HoughSegment test")
	}

	if dest.Rows() != 5 {
		t.Errorf("Invalid HoughSegment test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1 {
		t.Errorf("Invalid HoughSegment test cols: %v", dest.Cols())
	}

	type point struct {
		X, Y int32
	}

	expected := map[point]point{
		{1, 21}:   {398, 21},
		{304, 21}: {10, 315},
	}

	actual := make(map[point]point)
	for i := 0; i < dest.Rows(); i += 4 {
		actual[point{dest.GetVeciAt(i, 0)[0], dest.GetVeciAt(i, 0)[1]}] =
			point{dest.GetVeciAt(i, 0)[2], dest.GetVeciAt(i, 0)[3]}
	}

	for k, v := range expected {
		verify.Values(t, fmt.Sprintf("%d %d", k.X, k.Y), actual[k], v)
	}
}

func TestHoughSegment_CalcWithStream(t *testing.T) {
	src := engine.IMRead("../../images/face-detect.jpg", engine.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in HoughSegment test")
	}
	defer src.Close()

	cimg, mimg, dimg := NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer mimg.Close()
	defer dimg.Close()

	stream := NewStream()
	defer stream.Close()

	canny := NewCannyEdgeDetector(50, 100)
	defer canny.Close()

	detector := NewHoughSegmentDetector(1, math.Pi/180, 150, 50)
	defer detector.Close()

	dest := engine.NewMat()
	defer dest.Close()

	cimg.UploadWithStream(src, stream)
	canny.DetectWithStream(cimg, &mimg, stream)
	detector.DetectWithStream(mimg, &dimg, stream)
	fimg := dimg.Reshape(0, dimg.Cols())
	defer fimg.Close()
	fimg.DownloadWithStream(&dest, stream)

	stream.WaitForCompletion()

	if dest.Empty() {
		t.Error("Empty HoughSegment test")
	}

	if dest.Rows() != 5 {
		t.Errorf("Invalid HoughSegment test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1 {
		t.Errorf("Invalid HoughSegment test cols: %v", dest.Cols())
	}

	type point struct {
		X, Y int32
	}

	expected := map[point]point{
		{1, 21}:   {398, 21},
		{304, 21}: {10, 315},
	}

	actual := make(map[point]point)
	for i := 0; i < dest.Rows(); i += 4 {
		actual[point{dest.GetVeciAt(i, 0)[0], dest.GetVeciAt(i, 0)[1]}] =
			point{dest.GetVeciAt(i, 0)[2], dest.GetVeciAt(i, 0)[3]}
	}

	for k, v := range expected {
		verify.Values(t, fmt.Sprintf("%d %d", k.X, k.Y), actual[k], v)
	}
}
