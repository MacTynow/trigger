FROM golang:${GOLANG_VERSION} as build
WORKDIR /go/src/github.com/mactynow/trigger
COPY cmd/ cmd/
RUN CGO_ENABLED=0 GO111MODULE=on go build -a -o trigger cmd/main.go

FROM alpine:latest
WORKDIR /root
COPY --from=build go/src/github.com/mactynow/trigger/trigger . 
