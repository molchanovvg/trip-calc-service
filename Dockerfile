FROM golang:latest

WORKDIR /go/src/trip-calc-service

#COPY go.mod go.sum ./
#RUN go mod download
# RUN go get github.com/go-redis/redis/v8

COPY . .
#COPY calc storage structures ./

RUN go build -o main.go .

EXPOSE 8080

CMD ["./main"]