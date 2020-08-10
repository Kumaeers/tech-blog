FROM golang:latest

RUN apt-get update && \
    apt-get install -y default-mysql-client

# Go Modulesを使うと明言する
ENV GO111MODULE="on"

# ホットリロード機能を付与
RUN go get github.com/pilu/fresh

# マイグレーションツールの導入
RUN go get -u bitbucket.org/liamstask/goose/cmd/goose
