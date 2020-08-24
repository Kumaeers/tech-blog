# 開発環境
FROM golang:1.14.4 as dev
WORKDIR /go/src
# Go Modulesを使うと明言する
ENV GO111MODULE="on"
# ホットリロード機能を付与
RUN go get github.com/pilu/fresh
# マイグレーションツールの導入
RUN go get -v github.com/rubenv/sql-migrate/...
