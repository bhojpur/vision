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

#include "cuda.h"

GpuMat GpuMat_New() {
    return new cv::cuda::GpuMat();
}

GpuMat GpuMat_NewFromMat(Mat mat) {
    return new cv::cuda::GpuMat(*mat);
}

GpuMat GpuMat_NewWithSize(int rows, int cols, int type) {
    return new cv::cuda::GpuMat(rows, cols, type);
}

void GpuMat_Upload(GpuMat m, Mat data, Stream s){
    if (s == NULL) {
        m->upload(*data);
        return;
    }
    m->upload(*data, *s);
}

void GpuMat_Download(GpuMat m, Mat dst, Stream s){
    if (s == NULL) {
        m->download(*dst);
        return;
    }
    m->download(*dst, *s);
}

int GpuMat_Empty(GpuMat m){
    return m->empty();
}

void GpuMat_Close(GpuMat m){
    delete m;
}

void PrintCudaDeviceInfo(int device){
    cv::cuda::printCudaDeviceInfo(device);
}

void PrintShortCudaDeviceInfo(int device){
    cv::cuda::printShortCudaDeviceInfo(device);
}

int GetCudaEnabledDeviceCount(){
    return cv::cuda::getCudaEnabledDeviceCount();
}

int GetCudaDevice() {
    return cv::cuda::getDevice();
}

void SetCudaDevice(int device) {
    cv::cuda::setDevice(device);
}

void ResetCudaDevice(){
    cv::cuda::resetDevice();
}

void GpuMat_ConvertTo(GpuMat m, GpuMat dst, int type, Stream s) {
    if (s == NULL) {
        m->convertTo(*dst, type);
        return;
    }
    m->convertTo(*dst, type, *s);
}

void GpuMat_CopyTo(GpuMat m, GpuMat dst, Stream s) {
    if (s == NULL) {
        m->copyTo(*dst);
        return;
    }
    m->copyTo(*dst, *s);
}

GpuMat GpuMat_Reshape(GpuMat m, int cn, int rows) {
    return new cv::cuda::GpuMat(m->reshape(cn, rows));
}

int GpuMat_Cols(GpuMat m) {
    return m->cols;
}

int GpuMat_Rows(GpuMat m) {
    return m->rows;
}

int GpuMat_Channels(GpuMat m) {
    return m->channels();
}

int GpuMat_Type(GpuMat m) {
    return m->type();
}

Stream Stream_New() {
    return new cv::cuda::Stream();
}

void Stream_Close(Stream s){
    delete s;
}

bool Stream_QueryIfComplete(Stream s) {
    return s->queryIfComplete();
}

void Stream_WaitForCompletion(Stream s) {
    s->waitForCompletion();
}