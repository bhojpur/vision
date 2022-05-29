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

#include "imgcodecs.h"

// Image
Mat Image_IMRead(const char* filename, int flags) {
    cv::Mat img = cv::imread(filename, flags);
    return new cv::Mat(img);
}


bool Image_IMWrite(const char* filename, Mat img) {
    return cv::imwrite(filename, *img);
}

bool Image_IMWrite_WithParams(const char* filename, Mat img, IntVector params) {
    std::vector<int> compression_params;

    for (int i = 0, *v = params.val; i < params.length; ++v, ++i) {
        compression_params.push_back(*v);
    }

    return cv::imwrite(filename, *img, compression_params);
}

void Image_IMEncode(const char* fileExt, Mat img, void* vector) {
    auto vectorPtr = reinterpret_cast<std::vector<uchar> *>(vector);
    cv::imencode(fileExt, *img, *vectorPtr);
}

void Image_IMEncode_WithParams(const char* fileExt, Mat img, IntVector params, void* vector) {
    auto vectorPtr = reinterpret_cast<std::vector<uchar> *>(vector);
    std::vector<int> compression_params;

    for (int i = 0, *v = params.val; i < params.length; ++v, ++i) {
        compression_params.push_back(*v);
    }

    cv::imencode(fileExt, *img, *vectorPtr, compression_params);
}

Mat Image_IMDecode(ByteArray buf, int flags) {
    std::vector<uchar> data(buf.data, buf.data + buf.length);
    cv::Mat img = cv::imdecode(data, flags);
    return new cv::Mat(img);
}