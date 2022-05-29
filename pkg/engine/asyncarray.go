//go:build openvino
// +build openvino

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
	"errors"
)

/*
#include <stdlib.h>
#include "dnn.h"
#include "asyncarray.h"
#include "core.h"
*/
import "C"

type AsyncArray struct {
	p C.AsyncArray
}

// NewAsyncArray returns a new empty AsyncArray.
func NewAsyncArray() AsyncArray {
	return newAsyncArray(C.AsyncArray_New())
}

// Ptr returns the AsyncArray's underlying object pointer.
func (a *AsyncArray) Ptr() C.AsyncArray {
	return a.p
}

// Get async returns the Mat
func (m *AsyncArray) Get(mat *Mat) error {
	result := C.AsyncArray_GetAsync(m.p, mat.p)
	err := C.GoString(result)

	if len(err) > 0 {
		return errors.New(err)
	}
	return nil
}

// newAsyncArray returns a new AsyncArray from a C AsyncArray
func newAsyncArray(p C.AsyncArray) AsyncArray {
	return AsyncArray{p: p}
}

// Close the AsyncArray object.
func (a *AsyncArray) Close() error {
	C.AsyncArray_Close(a.p)
	a.p = nil
	return nil
}
