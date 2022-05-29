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
	"testing"
)

func TestSVDCompute(t *testing.T) {
	var resultW = []float32{6.167493, 3.8214223}
	var resultU = []float32{-0.1346676, -0.99089086, 0.9908908, -0.1346676}
	var resultVt = []float32{0.01964448, 0.999807, -0.999807, 0.01964448}

	checkFunc := func(a []float32, b []float32) bool {
		if len(a) != len(b) {
			return false
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	src := NewMatWithSize(2, 2, MatTypeCV32F)
	src.SetFloatAt(0, 0, 3.76956568)
	src.SetFloatAt(0, 1, -0.90478725)
	src.SetFloatAt(1, 0, 0.634576)
	src.SetFloatAt(1, 1, 6.10002347)
	defer src.Close()

	w := NewMat()
	defer w.Close()

	u := NewMat()
	defer u.Close()

	vt := NewMat()
	defer vt.Close()

	SVDCompute(src, &w, &u, &vt)

	dataW, err := w.DataPtrFloat32()
	if err != nil {
		t.Error(err)
	}

	if !checkFunc(resultW, dataW) {
		t.Error("w value is incorrect")
	}

	dataU, err := u.DataPtrFloat32()
	if err != nil {
		t.Error(err)
	}

	if !checkFunc(resultU, dataU) {
		t.Error("u value is incorrect")
	}

	dataVt, err := vt.DataPtrFloat32()
	if err != nil {
		t.Error(err)
	}

	if !checkFunc(resultVt, dataVt) {
		t.Error("vt value is incorrect")
	}
}
