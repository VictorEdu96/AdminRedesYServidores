FROM golang:latest

WORKDIR /go/src/app
COPY . .

EXPOSE 80
RUN go build -o server

CMD ["./server"]
