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
//
// It detects motion using a delta threshold from the first frame, and then
// finds contours to determine where the object is located.
//
// How to run:
//
// 		go run ./internal/motion/main.go 0
//

import (
	"fmt"
	"image"
	"image/color"
	"os"

	engine "github.com/bhojpur/vision/pkg/engine"
)

const MinimumArea = 3000

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tvision [camera ID]")
		return
	}

	// parse args
	deviceID := os.Args[1]

	webcam, err := engine.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := engine.NewWindow("Bhojpur Vision - Motion Window")
	defer window.Close()

	img := engine.NewMat()
	defer img.Close()

	imgDelta := engine.NewMat()
	defer imgDelta.Close()

	imgThresh := engine.NewMat()
	defer imgThresh.Close()

	mog2 := engine.NewBackgroundSubtractorMOG2()
	defer mog2.Close()

	status := "Ready"

	fmt.Printf("Start reading device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		status = "Ready"
		statusColor := color.RGBA{0, 255, 0, 0}

		// first phase of cleaning up image, obtain foreground only
		mog2.Apply(img, &imgDelta)

		// remaining cleanup of the image to use for finding contours.
		// first use threshold
		engine.Threshold(imgDelta, &imgThresh, 25, 255, engine.ThresholdBinary)

		// then dilate
		kernel := engine.GetStructuringElement(engine.MorphRect, image.Pt(3, 3))
		engine.Dilate(imgThresh, &imgThresh, kernel)
		kernel.Close()

		// now find contours
		contours := engine.FindContours(imgThresh, engine.RetrievalExternal, engine.ChainApproxSimple)

		for i := 0; i < contours.Size(); i++ {
			area := engine.ContourArea(contours.At(i))
			if area < MinimumArea {
				continue
			}

			status = "Motion detected"
			statusColor = color.RGBA{255, 0, 0, 0}
			engine.DrawContours(&img, contours, i, statusColor, 2)

			rect := engine.BoundingRect(contours.At(i))
			engine.Rectangle(&img, rect, color.RGBA{0, 0, 255, 0}, 2)
		}

		contours.Close()

		engine.PutText(&img, status, image.Pt(10, 20), engine.FontHersheyPlain, 1.2, statusColor, 2)

		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}
}
