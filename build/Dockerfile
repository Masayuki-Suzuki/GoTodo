FROM golang:1.19.2-alpine
RUN apk update && apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app

RUN go install github.com/cosmtrek/air@v1.29.0

# air -c [tomlファイル名] // 設定ファイルを指定してair実行(WORKDIRに.air.tomlを配置しておくこと)
CMD ["air", "-c", ".air.toml"]
