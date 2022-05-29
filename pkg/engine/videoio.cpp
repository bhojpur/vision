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

#include "videoio.h"

// VideoWriter
VideoCapture VideoCapture_New() {
    return new cv::VideoCapture();
}

void VideoCapture_Close(VideoCapture v) {
    delete v;
}

bool VideoCapture_Open(VideoCapture v, const char* uri) {
    return v->open(uri);
}

bool VideoCapture_OpenWithAPI(VideoCapture v, const char* uri, int apiPreference) {
    return v->open(uri, apiPreference);
}

bool VideoCapture_OpenDevice(VideoCapture v, int device) {
    return v->open(device);
}

bool VideoCapture_OpenDeviceWithAPI(VideoCapture v, int device, int apiPreference) {
    return v->open(device, apiPreference);
}

void VideoCapture_Set(VideoCapture v, int prop, double param) {
    v->set(prop, param);
}

double VideoCapture_Get(VideoCapture v, int prop) {
    return v->get(prop);
}

int VideoCapture_IsOpened(VideoCapture v) {
    return v->isOpened();
}

int VideoCapture_Read(VideoCapture v, Mat buf) {
    return v->read(*buf);
}

void VideoCapture_Grab(VideoCapture v, int skip) {
    for (int i = 0; i < skip; i++) {
        v->grab();
    }
}

// VideoWriter
VideoWriter VideoWriter_New() {
    return new cv::VideoWriter();
}

void VideoWriter_Close(VideoWriter vw) {
    delete vw;
}

void VideoWriter_Open(VideoWriter vw, const char* name, const char* codec, double fps, int width,
                      int height, bool isColor) {
    int codecCode = cv::VideoWriter::fourcc(codec[0], codec[1], codec[2], codec[3]);
    vw->open(name, codecCode, fps, cv::Size(width, height), isColor);
}

int VideoWriter_IsOpened(VideoWriter vw) {
    return vw->isOpened();
}

void VideoWriter_Write(VideoWriter vw, Mat img) {
    *vw << *img;
}