# Development Stage
FROM golang:1.24-alpine AS development

WORKDIR /app

# 必要なツールをインストール
RUN go install github.com/air-verse/air@latest

# モジュールファイルをコピーして依存関係を準備
COPY go.mod go.sum ./
RUN go mod tidy

# ソースコードをコピー
COPY . ./

# 開発用のコマンド
CMD ["air"]

# Build Stage
FROM golang:1.23-alpine AS build

WORKDIR /app

# モジュールファイルをコピーして依存関係を準備
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY . ./
RUN go build -o main .

# Production Stage
FROM alpine:latest AS production

WORKDIR /app

# 必要な実行ファイルだけをコピー
COPY --from=build /app/main .

# 必要なポートを公開
EXPOSE 8080

# 本番用のコマンド
CMD ["./main"]
