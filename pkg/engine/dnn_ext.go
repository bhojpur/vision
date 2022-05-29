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
)

// FP16BlobFromImage is an extended helper function to convert an Image to a half-float blob, as used by
// the Movidius Neural Compute Stick.
func FP16BlobFromImage(img Mat, scaleFactor float32, size image.Point, mean float32,
	swapRB bool, crop bool) []byte {

	// resizes image so it maintains aspect ratio
	width := float32(img.Cols())
	height := float32(img.Rows())

	square := NewMatWithSize(size.Y, size.X, img.Type())
	defer square.Close()

	maxDim := height
	var scale float32 = 1.0
	if width > height {
		maxDim = width
		scale = float32(size.X) / float32(maxDim)
	}
	if width < height {
		scale = float32(size.Y) / float32(maxDim)
	}

	var roi image.Rectangle
	if width >= height {
		roi.Min.X = 0
		roi.Min.Y = int(float32(size.Y)-height*scale) / 2
		roi.Max.X = size.X
		roi.Max.Y = int(height * scale)
	} else {
		roi.Min.X = int(float32(size.X)-width*scale) / 2
		roi.Min.Y = 0
		roi.Max.X = int(width * scale)
		roi.Max.Y = size.Y
	}

	Resize(img, &square, roi.Max, 0, 0, InterpolationDefault)

	if swapRB {
		CvtColor(square, &square, ColorBGRToRGB)
	}

	fp32Image := NewMat()
	defer fp32Image.Close()

	square.ConvertTo(&fp32Image, MatTypeCV32F)

	if mean != 0 {
		// subtract mean
		fp32Image.SubtractFloat(mean)
	}

	if scaleFactor != 1.0 {
		// multiply by scale factor
		fp32Image.MultiplyFloat(scaleFactor)
	}

	fp16Blob := fp32Image.ConvertFp16()
	defer fp16Blob.Close()

	return fp16Blob.ToBytes()
}
