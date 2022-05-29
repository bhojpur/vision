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

#ifndef _OPENCV3_HIGHGUI_H_
#define _OPENCV3_HIGHGUI_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

// Window
void Window_New(const char* winname, int flags);
void Window_Close(const char* winname);
void Window_IMShow(const char* winname, Mat mat);
double Window_GetProperty(const char* winname, int flag);
void Window_SetProperty(const char* winname, int flag, double value);
void Window_SetTitle(const char* winname, const char* title);
int Window_WaitKey(int);
void Window_Move(const char* winname, int x, int y);
void Window_Resize(const char* winname, int width, int height);
struct Rect Window_SelectROI(const char* winname, Mat img);
struct Rects Window_SelectROIs(const char* winname, Mat img);

// Trackbar
void Trackbar_Create(const char* winname, const char* trackname, int max);
void Trackbar_CreateWithValue(const char* winname, const char* trackname, int* value, int max);
int Trackbar_GetPos(const char* winname, const char* trackname);
void Trackbar_SetPos(const char* winname, const char* trackname, int pos);
void Trackbar_SetMin(const char* winname, const char* trackname, int pos);
void Trackbar_SetMax(const char* winname, const char* trackname, int pos);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_HIGHGUI_H_