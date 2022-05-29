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

func TestNewGpuMat(t *testing.T) {
	mat := NewGpuMat()
	defer mat.Close()

	if !mat.Empty() {
		t.Error("New GpuMat should be empty")
	}
}

func TestNewGpuMatFromMat(t *testing.T) {
	mat := engine.NewMat()
	defer mat.Close()

	gpumat := NewGpuMatFromMat(mat)
	defer gpumat.Close()

	if !gpumat.Empty() {
		t.Error("New GpuMat should be empty")
	}
}

func TestNewGpuMatFromMatWithSize(t *testing.T) {
	mat := engine.NewMatWithSize(100, 200, engine.MatTypeCV32FC4)
	defer mat.Close()

	gpumat := NewGpuMatFromMat(mat)
	defer gpumat.Close()

	if gpumat.Empty() {
		t.Error("New GpuMat should be not empty")
	}

	if gpumat.Rows() != 100 {
		t.Error("incorrect number of rows for GpuMat")
	}

	if gpumat.Cols() != 200 {
		t.Error("incorrect number of cols for GpuMat")
	}

	if gpumat.Type() != engine.MatTypeCV32FC4 {
		t.Error("incorrect type for GpuMat")
	}
}

func TestNewGpuMatWithSize(t *testing.T) {
	gpumat := NewGpuMatWithSize(100, 200, engine.MatTypeCV32FC4)
	defer gpumat.Close()

	if gpumat.Empty() {
		t.Error("New GpuMat should be not empty")
	}

	if gpumat.Rows() != 100 {
		t.Error("incorrect number of rows for GpuMat")
	}

	if gpumat.Cols() != 200 {
		t.Error("incorrect number of cols for GpuMat")
	}

	if gpumat.Type() != engine.MatTypeCV32FC4 {
		t.Error("incorrect type for GpuMat")
	}
}

func TestGetCudaEnabledDeviceCount(t *testing.T) {
	if GetCudaEnabledDeviceCount() < 1 {
		t.Fatal("expected atleast one cuda enabled device")
	}
}
