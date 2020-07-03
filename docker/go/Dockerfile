FROM golang:1.14-alpine

RUN apk update && \
    apk add mysql-client

# Go Modulesを使うと明言する
ENV GO111MODULE="on"

# ホットリロード機能を付与
RUN go get github.com/pilu/fresh
