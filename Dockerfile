# ベースイメージを指定
FROM golang:1.23.2

# 必要なパッケージをインストール
RUN apt-get update && apt-get install -y \
    default-mysql-client \
    locales && \
    echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen && \
    update-locale LANG=ja_JP.UTF-8

# 環境変数を設定
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8

# 作業ディレクトリを設定
WORKDIR /app

# Air をインストール
RUN go install github.com/cosmtrek/air@v1.40.4

# Goose をインストール
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.22.1

# ローカルのソースコードをコンテナにコピー
COPY . .

# 必要なパッケージをインストール
RUN go mod download

# ポートを公開
EXPOSE 8080

# Air を使ってアプリケーションをホットリロード
CMD ["air"]
