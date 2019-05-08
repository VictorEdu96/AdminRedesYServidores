FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -d -v github.com/chzyer/readline
RUN go build -o client
CMD ["./client"]
CMD ["conn", "localhost:80"]
