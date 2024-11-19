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

# direnv をインストール
RUN curl -sfL https://direnv.net/install.sh | bash

# direnv を PATH に追加
ENV PATH="$PATH:/usr/local/bin"

# direnv フックをシェル設定に追加
RUN echo 'eval "$(direnv hook bash)"' >> /etc/bash.bashrc

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# ローカルのソースコードをコンテナにコピー
COPY . .

# 必要なパッケージをインストール
RUN go mod download

# .envrc の許可を自動化
RUN direnv allow || true

# アプリケーションをビルド
RUN go build -o main .

# コンテナのポートを公開
EXPOSE 8080

# アプリケーションを実行
CMD ["go", "run", "main.go"]
