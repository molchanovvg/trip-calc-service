FROM golang:latest

WORKDIR /go/src/trip-calc-service

#COPY go.mod go.sum ./

#RUN go get github.com/go-redis/redis

COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...

RUN go build -o main .

ENV REDIS_URL=localhost:6379
ENV SERVICE_PORT=8080
ENV CALC_ROUTE_URL=http://router.project-osrm.org/route/v1/driving/

EXPOSE 8080

CMD ["go", "run", "main.go"]