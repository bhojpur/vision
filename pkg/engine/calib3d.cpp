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

#include "calib3d.h"

void Fisheye_UndistortImage(Mat distorted, Mat undistorted, Mat k, Mat d) {
    cv::fisheye::undistortImage(*distorted, *undistorted, *k, *d);
}

void Fisheye_UndistortImageWithParams(Mat distorted, Mat undistorted, Mat k, Mat d, Mat knew, Size size) {
    cv::Size sz(size.width, size.height);
    cv::fisheye::undistortImage(*distorted, *undistorted, *k, *d, *knew, sz);
}

void Fisheye_UndistortPoints(Mat distorted, Mat undistorted, Mat k, Mat d, Mat r, Mat p) {
    cv::fisheye::undistortPoints(*distorted, *undistorted, *k, *d, *r, *p);
}

void Fisheye_EstimateNewCameraMatrixForUndistortRectify(Mat k, Mat d, Size imgSize, Mat r, Mat p, double balance, Size newSize, double fovScale) {
    cv::Size newSz(newSize.width, newSize.height);
    cv::Size imgSz(imgSize.width, imgSize.height);
    cv::fisheye::estimateNewCameraMatrixForUndistortRectify(*k, *d, imgSz, *r, *p, balance, newSz, fovScale);
}

void InitUndistortRectifyMap(Mat cameraMatrix,Mat distCoeffs,Mat r,Mat newCameraMatrix,Size size,int m1type,Mat map1,Mat map2) {
    cv::Size sz(size.width, size.height);
    cv::initUndistortRectifyMap(*cameraMatrix,*distCoeffs,*r,*newCameraMatrix,sz,m1type,*map1,*map2);
}

Mat GetOptimalNewCameraMatrixWithParams(Mat cameraMatrix,Mat distCoeffs,Size size,double alpha,Size newImgSize,Rect* validPixROI,bool centerPrincipalPoint) {
    cv::Size sz(size.width, size.height);
    cv::Size newSize(newImgSize.width, newImgSize.height);
    cv::Rect rect(validPixROI->x,validPixROI->y,validPixROI->width,validPixROI->height);
    cv::Mat* mat = new cv::Mat(cv::getOptimalNewCameraMatrix(*cameraMatrix,*distCoeffs,sz,alpha,newSize,&rect,centerPrincipalPoint));
    validPixROI->x = rect.x;
    validPixROI->y = rect.y;
    validPixROI->width = rect.width;
    validPixROI->height = rect.height;
    return mat;
}

double CalibrateCamera(Points3fVector objectPoints, Points2fVector imagePoints, Size imageSize, Mat cameraMatrix, Mat distCoeffs, Mat rvecs, Mat tvecs, int flag) {
    return cv::calibrateCamera(*objectPoints, *imagePoints, cv::Size(imageSize.width, imageSize.height), *cameraMatrix, *distCoeffs, *rvecs, *tvecs, flag);
}

void Undistort(Mat src, Mat dst, Mat cameraMatrix, Mat distCoeffs, Mat newCameraMatrix) {
    cv::undistort(*src, *dst, *cameraMatrix, *distCoeffs, *newCameraMatrix);
}

void UndistortPoints(Mat distorted, Mat undistorted, Mat k, Mat d, Mat r, Mat p) {
    cv::undistortPoints(*distorted, *undistorted, *k, *d, *r, *p);
}

bool FindChessboardCorners(Mat image, Size patternSize, Mat corners, int flags) {
    cv::Size sz(patternSize.width, patternSize.height);
    return cv::findChessboardCorners(*image, sz, *corners, flags);
}

bool FindChessboardCornersSB(Mat image, Size patternSize, Mat corners, int flags) {
    cv::Size sz(patternSize.width, patternSize.height);
    return cv::findChessboardCornersSB(*image, sz, *corners, flags);
}

bool FindChessboardCornersSBWithMeta(Mat image, Size patternSize, Mat corners, int flags, Mat meta) {
    cv::Size sz(patternSize.width, patternSize.height);
    return cv::findChessboardCornersSB(*image, sz, *corners, flags, *meta);
}

void DrawChessboardCorners(Mat image, Size patternSize, Mat corners, bool patternWasFound) {
    cv::Size sz(patternSize.width, patternSize.height);
    cv::drawChessboardCorners(*image, sz, *corners, patternWasFound);
}

Mat EstimateAffinePartial2D(Point2fVector from, Point2fVector to) {
    return new cv::Mat(cv::estimateAffinePartial2D(*from, *to));
}

Mat EstimateAffine2D(Point2fVector from, Point2fVector to) {
    return new cv::Mat(cv::estimateAffine2D(*from, *to));
}

Mat EstimateAffine2DWithParams(Point2fVector from, Point2fVector to, Mat inliers, int method, double ransacReprojThreshold, size_t maxIters, double confidence, size_t refineIters) {
    return new cv::Mat(cv::estimateAffine2D(*from, *to, *inliers, method, ransacReprojThreshold, maxIters, confidence, refineIters));
}