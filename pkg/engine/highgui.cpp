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

#include "highgui_vision.h"

// Window
void Window_New(const char* winname, int flags) {
    cv::namedWindow(winname, flags);
}

void Window_Close(const char* winname) {
    cv::destroyWindow(winname);
}

void Window_IMShow(const char* winname, Mat mat) {
    cv::imshow(winname, *mat);
}

double Window_GetProperty(const char* winname, int flag) {
    return cv::getWindowProperty(winname, flag);
}

void Window_SetProperty(const char* winname, int flag, double value) {
    cv::setWindowProperty(winname, flag, value);
}

void Window_SetTitle(const char* winname, const char* title) {
    cv::setWindowTitle(winname, title);
}

int Window_WaitKey(int delay = 0) {
    return cv::waitKey(delay);
}

void Window_Move(const char* winname, int x, int y) {
    cv::moveWindow(winname, x, y);
}

void Window_Resize(const char* winname, int width, int height) {
    cv::resizeWindow(winname, width, height);
}

struct Rect Window_SelectROI(const char* winname, Mat img) {
    cv::Rect bRect = cv::selectROI(winname, *img);
    Rect r = {bRect.x, bRect.y, bRect.width, bRect.height};
    return r;
}

struct Rects Window_SelectROIs(const char* winname, Mat img) {
    std::vector<cv::Rect> rois;
    cv::selectROIs(winname, *img, rois);
    Rect* rects = new Rect[rois.size()];

    for (size_t i = 0; i < rois.size(); ++i) {
        Rect r = {rois[i].x, rois[i].y, rois[i].width, rois[i].height};
        rects[i] = r;
    }

    Rects ret = {rects, (int)rois.size()};
    return ret;
}

// Trackbar
void Trackbar_Create(const char* winname, const char* trackname, int max) {
    cv::createTrackbar(trackname, winname, NULL, max);
}

void Trackbar_CreateWithValue(const char* winname, const char* trackname, int* value, int max) {
    cv::createTrackbar(trackname, winname, value, max);
}

int Trackbar_GetPos(const char* winname, const char* trackname) {
    return cv::getTrackbarPos(trackname, winname);
}

void Trackbar_SetPos(const char* winname, const char* trackname, int pos) {
    cv::setTrackbarPos(trackname, winname, pos);
}

void Trackbar_SetMin(const char* winname, const char* trackname, int pos) {
    cv::setTrackbarMin(trackname, winname, pos);
}

void Trackbar_SetMax(const char* winname, const char* trackname, int pos) {
    cv::setTrackbarMax(trackname, winname, pos);
}