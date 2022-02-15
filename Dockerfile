# goバージョン
FROM golang:1.17-alpine
# アップデートとgitのインストール！！
RUN apk add --update &&  apk add git
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
