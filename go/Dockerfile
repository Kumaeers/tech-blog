# ビルドステージ
FROM golang:1.14.4 as build-stage
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

# 本番環境
FROM scratch as production-stage
COPY --from=build-stage /go/bin/main /app/main
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8082
ENTRYPOINT ["/app/main"]
