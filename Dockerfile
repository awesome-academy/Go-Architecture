FROM golang:1.21-alpine

WORKDIR /usr/src/app

COPY . .

RUN go install github.com/cosmtrek/air@latest
RUN go get gorm.io/gorm
RUN go get github.com/gin-gonic/gin
RUN go get github.com/golang-jwt/jwt
RUN go get github.com/lib/pq
RUN go mod tidy
