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

func (c CalibFlag) String() string {
	switch c {
	case CalibUseIntrinsicGuess:
		return "calib-use-intrinsec-guess"
	case CalibRecomputeExtrinsic:
		return "calib-recompute-extrinsic"
	case CalibCheckCond:
		return "calib-check-cond"
	case CalibFixSkew:
		return "calib-fix-skew"
	case CalibFixK1:
		return "calib-fix-k1"
	case CalibFixK2:
		return "calib-fix-k2"
	case CalibFixK3:
		return "calib-fix-k3"
	case CalibFixK4:
		return "calib-fix-k4"
	case CalibFixIntrinsic:
		return "calib-fix-intrinsic"
	case CalibFixPrincipalPoint:
		return "calib-fix-principal-point"
	}
	return ""
}

func (c CalibCBFlag) String() string {
	switch c {
	case CalibCBAdaptiveThresh:
		return "calib-cb-adaptive-thresh"
	case CalibCBNormalizeImage:
		return "calib-cb-normalize-image"
	case CalibCBFilterQuads:
		return "calib-cb-filter-quads"
	case CalibCBFastCheck:
		return "calib-cb-fast-check"
	case CalibCBExhaustive:
		return "calib-cb-exhaustive"
	case CalibCBAccuracy:
		return "calib-cb-accuracy"
	case CalibCBLarger:
		return "calib-cb-larger"
	case CalibCBMarker:
		return "calib-cb-marker"
	}
	return ""
}
