# ベースイメージ
FROM golang:1.24-alpine

# 必要なツールをインストール
RUN apk add --no-cache bash git

# 作業ディレクトリの作成
WORKDIR /app

# Goの依存ファイルをコピーしてキャッシュを活かす
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをすべてコピー
COPY . .

# アプリケーションをビルド（main.go が cmd/ にある場合）
RUN go build -o main ./cmd/main.go

# ポートを公開
EXPOSE 8080

# wait-for-it.sh をコピーして実行権限付与
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# アプリを起動（MySQLが起動してから）
CMD ["./wait-for-it.sh", "db:3306", "--", "./main"]