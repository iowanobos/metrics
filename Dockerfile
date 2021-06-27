FROM golang:1.15.11-alpine3.12 AS builder

RUN apk update
RUN apk --no-cache add git make ffmpeg

WORKDIR /go/src/app
COPY . .
RUN make build

CMD ["./bin/metrics"]
