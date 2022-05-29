package engine

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
)

func TestMOG2(t *testing.T) {
	img := IMRead("../../images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in MOG2 test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	mog2 := NewBackgroundSubtractorMOG2()
	defer mog2.Close()

	mog2.Apply(img, &dst)

	if dst.Empty() {
		t.Error("Error in TestMOG2 test")
	}
}

func TestMOG2WithParams(t *testing.T) {
	img := IMRead("../../images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in MOG2 test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	mog2 := NewBackgroundSubtractorMOG2WithParams(250, 8, false)
	defer mog2.Close()

	mog2.Apply(img, &dst)

	if dst.Empty() {
		t.Error("Error in TestMOG2WithParams test")
	}
}

func TestKNN(t *testing.T) {
	img := IMRead("../../images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in KNN test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	knn := NewBackgroundSubtractorKNN()
	defer knn.Close()

	knn.Apply(img, &dst)

	if dst.Empty() {
		t.Error("Error in TestKNN test")
	}
}

func TestKNNWithParams(t *testing.T) {
	img := IMRead("../../images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in KNN test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	knn := NewBackgroundSubtractorKNNWithParams(250, 200, false)
	defer knn.Close()

	knn.Apply(img, &dst)

	if dst.Empty() {
		t.Error("Error in TestKNNWithParams test")
	}
}

func TestCalcOpticalFlowFarneback(t *testing.T) {
	img1 := IMRead("../../images/face.jpg", IMReadColor)
	if img1.Empty() {
		t.Error("Invalid Mat in CalcOpticalFlowFarneback test")
	}
	defer img1.Close()

	dest := NewMat()
	defer dest.Close()

	CvtColor(img1, &dest, ColorBGRAToGray)

	img2 := dest.Clone()
	defer img2.Close()

	flow := NewMat()
	defer flow.Close()

	CalcOpticalFlowFarneback(dest, img2, &flow, 0.4, 1, 12, 2, 8, 1.2, 0)

	if flow.Empty() {
		t.Error("Error in CalcOpticalFlowFarneback test")
	}
	if flow.Rows() != 480 {
		t.Errorf("Invalid CalcOpticalFlowFarneback test rows: %v", flow.Rows())
	}
	if flow.Cols() != 640 {
		t.Errorf("Invalid CalcOpticalFlowFarneback test cols: %v", flow.Cols())
	}
}

func TestCalcOpticalFlowPyrLK(t *testing.T) {
	img1 := IMRead("../../images/face.jpg", IMReadColor)
	if img1.Empty() {
		t.Error("Invalid Mat in CalcOpticalFlowPyrLK test")
	}
	defer img1.Close()

	dest := NewMat()
	defer dest.Close()

	CvtColor(img1, &dest, ColorBGRAToGray)

	img2 := dest.Clone()
	defer img2.Close()

	prevPts := NewMat()
	defer prevPts.Close()

	nextPts := NewMat()
	defer nextPts.Close()

	status := NewMat()
	defer status.Close()

	err := NewMat()
	defer err.Close()

	corners := NewMat()
	defer corners.Close()

	GoodFeaturesToTrack(dest, &corners, 500, 0.01, 10)
	tc := NewTermCriteria(Count|EPS, 20, 0.03)
	CornerSubPix(dest, &corners, image.Pt(10, 10), image.Pt(-1, -1), tc)

	CalcOpticalFlowPyrLK(dest, img2, corners, nextPts, &status, &err)

	if status.Empty() {
		t.Error("Error in CalcOpticalFlowPyrLK test")
	}
	if status.Rows() != 323 {
		t.Errorf("Invalid CalcOpticalFlowPyrLK test rows: %v", status.Rows())
	}
	if status.Cols() != 1 {
		t.Errorf("Invalid CalcOpticalFlowPyrLK test cols: %v", status.Cols())
	}
}

func TestCalcOpticalFlowPyrLKWithParams(t *testing.T) {
	img1 := IMRead("../../images/face.jpg", IMReadColor)
	if img1.Empty() {
		t.Error("Invalid Mat in CalcOpticalFlowPyrLK test")
	}
	defer img1.Close()

	dest := NewMat()
	defer dest.Close()

	CvtColor(img1, &dest, ColorBGRAToGray)

	img2 := dest.Clone()
	defer img2.Close()

	prevPts := NewMat()
	defer prevPts.Close()

	nextPts := NewMat()
	defer nextPts.Close()

	status := NewMat()
	defer status.Close()

	err := NewMat()
	defer err.Close()

	corners := NewMat()
	defer corners.Close()

	GoodFeaturesToTrack(dest, &corners, 500, 0.01, 10)
	tc := NewTermCriteria(Count|EPS, 30, 0.03)
	CornerSubPix(dest, &corners, image.Pt(10, 10), image.Pt(-1, -1), tc)

	CalcOpticalFlowPyrLKWithParams(dest, img2, corners, nextPts, &status, &err, image.Pt(21, 21), 3, tc, 0, 0.0001)

	if status.Empty() {
		t.Error("Error in CalcOpticalFlowPyrLK test")
	}
	if status.Rows() != 323 {
		t.Errorf("Invalid CalcOpticalFlowPyrLK test rows: %v", status.Rows())
	}
	if status.Cols() != 1 {
		t.Errorf("Invalid CalcOpticalFlowPyrLK test cols: %v", status.Cols())
	}
}

func BaseTestTracker(t *testing.T, tracker Tracker, name string) {
	if tracker == nil {
		t.Error("TestTracker " + name + " should not be nil")
	}

	img := IMRead("./images/face.jpg", 1)
	if img.Empty() {
		t.Error("TestTracker " + name + " input img failed to load")
	}
	defer img.Close()

	rect := image.Rect(250, 150, 250+200, 150+250)
	init := tracker.Init(img, rect)
	if !init {
		t.Error("TestTracker " + name + " failed in Init")
	}

	_, ok := tracker.Update(img)
	if !ok {
		t.Error("TestTracker " + name + " lost object in Update")
	}
}

func TestSingleTrackers(t *testing.T) {
	tab := []struct {
		name    string
		tracker Tracker
	}{
		{"MIL", NewTrackerMIL()},
		// {"GOTURN", NewTrackerGOTURN()},
	}

	for _, test := range tab {
		func() {
			defer test.tracker.Close()
			BaseTestTracker(t, test.tracker, test.name)
		}()
	}
}
