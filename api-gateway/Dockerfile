FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o api-gateway

EXPOSE 4444 

CMD [ "./api-gateway" ]
