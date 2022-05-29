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

#ifndef _OPENCV3_OBJDETECT_H_
#define _OPENCV3_OBJDETECT_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

#ifdef __cplusplus
typedef cv::CascadeClassifier* CascadeClassifier;
typedef cv::HOGDescriptor* HOGDescriptor;
typedef cv::QRCodeDetector* QRCodeDetector;
#else
typedef void* CascadeClassifier;
typedef void* HOGDescriptor;
typedef void* QRCodeDetector;
#endif

// CascadeClassifier
CascadeClassifier CascadeClassifier_New();
void CascadeClassifier_Close(CascadeClassifier cs);
int CascadeClassifier_Load(CascadeClassifier cs, const char* name);
struct Rects CascadeClassifier_DetectMultiScale(CascadeClassifier cs, Mat img);
struct Rects CascadeClassifier_DetectMultiScaleWithParams(CascadeClassifier cs, Mat img,
        double scale, int minNeighbors, int flags, Size minSize, Size maxSize);

HOGDescriptor HOGDescriptor_New();
void HOGDescriptor_Close(HOGDescriptor hog);
int HOGDescriptor_Load(HOGDescriptor hog, const char* name);
struct Rects HOGDescriptor_DetectMultiScale(HOGDescriptor hog, Mat img);
struct Rects HOGDescriptor_DetectMultiScaleWithParams(HOGDescriptor hog, Mat img,
        double hitThresh, Size winStride, Size padding, double scale, double finalThreshold,
        bool useMeanshiftGrouping);
Mat HOG_GetDefaultPeopleDetector();
void HOGDescriptor_SetSVMDetector(HOGDescriptor hog, Mat det);

struct Rects GroupRectangles(struct Rects rects, int groupThreshold, double eps);

QRCodeDetector QRCodeDetector_New();
const char* QRCodeDetector_DetectAndDecode(QRCodeDetector qr, Mat input,Mat points,Mat straight_qrcode);
bool QRCodeDetector_Detect(QRCodeDetector qr, Mat input,Mat points);
const char* QRCodeDetector_Decode(QRCodeDetector qr, Mat input,Mat inputPoints,Mat straight_qrcode);
void QRCodeDetector_Close(QRCodeDetector qr);
bool QRCodeDetector_DetectMulti(QRCodeDetector qr, Mat input, Mat points);
bool QRCodeDetector_DetectAndDecodeMulti(QRCodeDetector qr, Mat input, CStrings* decoded ,Mat points, struct Mats* mats);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_OBJDETECT_H_