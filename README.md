# Bhojpur Vision - Processing Engine

The `Bhojpur Vision` is an advanced visual data processing engine applied within the
[Bhojpur.NET Platform](https://github.com/bhojpur/platform/) for delivery of intelligent
`applications` or `services`.

## Build Source Code

You need to install `Go` v1.16 at least. Also, [CUDA](https://docs.nvidia.com/cuda/index.html)
is required for faster GPU processing.

### Object Detection Models

Firstly, you need to do the following to achieve automatic object detection using ML. The
[PASCAL VOC](http://host.robots.ox.ac.uk/pascal/VOC/) (in .xml file format) and
[DarkNet](https://github.com/pjreddie/darknet) (in .txt file format) could be used for annotations.

- Data Acquisition
- Data Preparation according to the [Yolo](https://pjreddie.com/darknet/yolo/)
- Load Dataset
- Train the dataset
- Obtain the model weights
- Test the model

Alternatively, you can use [Yolo](https://pjreddie.com/darknet/yolo/) v5 pre-trained models.
Please run the following commands to download sample object detection models.

```bash
./getModels.sh
```

### Simple Examples

```bash
go run ./internal/image/main.go
go run ./internal/webcam/main.go
```
