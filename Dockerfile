# to build this docker image:
#   docker build .
FROM bhojpur/opencv:4.5.5

ENV GOPATH /go

COPY . /go/src/github.com/bhojpur/vision/

WORKDIR /go/src/github.com/bhojpur/vision
RUN go build -tags example -o /build/vision_version -i ./cmd/version/

CMD ["/build/vision_version"]