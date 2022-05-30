package main

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
	"path"

	engine "github.com/bhojpur/vision/pkg/engine"
	yolov5 "github.com/bhojpur/vision/pkg/yolo"
	log "github.com/sirupsen/logrus"
)

var (
	yolov5Model   = path.Join("./pkg/", "data/yolov5/yolov5s.onnx")
	cocoNamesPath = path.Join("./pkg/", "data/yolov5/coco.names")
)

func main() {
	yolonet, err := yolov5.NewNet(yolov5Model, cocoNamesPath)
	if err != nil {
		log.WithError(err).Fatal("unable to create yolo net")
	}

	// Gracefully close the net when the program is done
	defer func() {
		err := yolonet.Close()
		if err != nil {
			log.WithError(err).Error("unable to gracefully close yolo net")
		}
	}()

	videoCapture, err := engine.OpenVideoCapture(0)
	if err != nil {
		log.WithError(err).Fatal("unable to start video capture")
	}

	window := engine.NewWindow("Bhojpur Vision - Results Window")
	defer func() {
		err := window.Close()
		if err != nil {
			log.WithError(err).Error("unable to close window")
		}
	}()

	frame := engine.NewMat()
	defer func() {
		err := frame.Close()
		if err != nil {
			log.WithError(err).Errorf("unable to close image")
		}
	}()

	for {
		if ok := videoCapture.Read(&frame); !ok {
			log.Error("unable to read videostram")
		}
		if frame.Empty() {
			continue
		}
		detections, err := yolonet.GetDetections(frame)
		if err != nil {
			log.WithError(err).Fatal("unable to retrieve predictions")
		}

		yolov5.DrawDetections(&frame, detections)

		window.IMShow(frame)
		window.WaitKey(1)
	}
}
