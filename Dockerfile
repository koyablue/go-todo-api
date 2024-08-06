# Goバージョン
FROM golang:1.19.1-alpine

# 必要なツールのインストール
RUN apk update && apk add --no-cache git

# ワーキングディレクトリの設定
WORKDIR /app

# モジュールファイルをコピー
COPY go.mod ./
COPY go.sum ./

# モジュールのダウンロード
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションのビルド
RUN go build -o /app/go_todo_api main.go

# 実行時のコマンド
CMD ["/app/go_todo_api"]
