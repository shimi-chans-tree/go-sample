FROM golang:1.17.7-alpine

RUN apk update &&  apk add git
RUN go get github.com/cosmtrek/air@v1.29.0


WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

# air -c [tomlファイル名] // 設定ファイルを指定してair実行(WORKDIRに.air.tomlを配置しておくこと)
CMD ["air", "-c", ".air.toml"]


