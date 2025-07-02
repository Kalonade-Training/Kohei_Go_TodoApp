# ベースイメージ
FROM golang:1.20-alpine

# 必要なツールのインストール
RUN apk add --no-cache bash git

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールキャッシュを利用して依存関係を取得
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# ポート番号の公開
EXPOSE 8080

# wait-for-it.sh をコピーして実行権限付与
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# MySQLの起動を待ってからアプリを起動
CMD ["./wait-for-it.sh", "db:3306", "--", "./main"]