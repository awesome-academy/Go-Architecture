FROM golang:1.21

RUN mkdir /app

ADD . /app

ADD .env /

WORKDIR /app

RUN go build -o main cmd/main.go

CMD ["/app/main"]
