FROM golang:1.18-alpine

RUN apk add build-base

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o scheduler-service

CMD [ "./scheduler-service" ]
