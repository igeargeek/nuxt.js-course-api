FROM golang:1.12

ENV GO111MODULE=on

WORKDIR /go/src/app

COPY . .

RUN go build -o main .
COPY go.mod ./go/src/app
COPY go.sum ./go/src/app

RUN go mod download

EXPOSE 8000

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o main ." -command="./main"