# FROM golang:latest
FROM golang:1.12-rc

RUN mkdir /twifkak-go
WORKDIR /twifkak-go

# This golang fork makes it possible to run WASM on Chrome for Android
# https://github.com/golang/go/compare/master...twifkak:small
RUN git clone https://github.com/olso/go

WORKDIR /twifkak-go/go
RUN git fetch --all
RUN git checkout small

WORKDIR /twifkak-go/go/src
RUN ./make.bash

ENV GOPATH /twifkak-go/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
