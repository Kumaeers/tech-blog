# 開発環境
FROM golang:1.14.4 as dev
WORKDIR /go/src
# Go Modulesを使うと明言する
ENV GO111MODULE="on"
# ホットリロード機能を付与
RUN go get github.com/pilu/fresh

# 本番ビルド用
FROM golang:1.14.4 as builder
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'

# 本番用
FROM scratch as runner
COPY --from=builder /go/bin/main /app/main
EXPOSE 8082
ENTRYPOINT ["/app/main"]
