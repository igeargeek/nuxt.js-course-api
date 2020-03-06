FROM golang:1.14

ENV GO111MODULE=on

WORKDIR /usr/local/go/src/app

RUN go get github.com/google/wire/cmd/wire \
    && go get github.com/githubnemo/CompileDaemon

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN wire

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o main ." -command="./main"

EXPOSE 8000