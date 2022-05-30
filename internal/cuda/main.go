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
	"fmt"
	"path"
	"time"

	engine "github.com/bhojpur/vision/pkg/engine"
	yolov5 "github.com/bhojpur/vision/pkg/yolo"
	log "github.com/sirupsen/logrus"
)

var (
	yolov5Model   = path.Join("./pkg/", "data/yolov5/yolov5s.onnx")
	cocoNamesPath = path.Join("./pkg/", "data/yolov5/coco.names")
)

func main() {
	conf := yolov5.DefaultConfig()
	conf.NetBackendType = engine.NetBackendCUDA
	conf.NetTargetType = engine.NetTargetCUDA

	yolonet, err := yolov5.NewNetWithConfig(yolov5Model, cocoNamesPath, conf)
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

	window := engine.NewWindow("Bhojpur Vision - Results Window")
	defer func() {
		err := window.Close()
		if err != nil {
			log.WithError(err).Error("unable to close window")
		}
	}()
	window.ResizeWindow(872, 585)
	orgFrame := engine.IMRead(path.Join("./pkg/", "data/example_images/bird.jpg"), engine.IMReadColor)
	defer func() {
		err := orgFrame.Close()
		if err != nil {
			log.WithError(err).Error("unable to close frame")
		}
	}()

	// Render example image at 50 frames a second
	ticker := time.NewTicker(time.Second / 50)
	go func() {
		for {
			<-ticker.C
			frame := orgFrame.Clone()
			detections, err := yolonet.GetDetections(frame)
			if err != nil {
				err = fmt.Errorf("%w %s", err, frame.Close())
				log.WithError(err).Fatal("unable to retrieve predictions")
				continue
			}

			yolov5.DrawDetections(&frame, detections)

			window.IMShow(frame)
			err = frame.Close()
			if err != nil {
				log.WithError(err).Error("unable to close frame")
			}
		}
	}()
	window.WaitKey(10000000000)
}
