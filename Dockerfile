FROM golang:1.20.5-alpine3.18
WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify


RUN apk add --no-cache make

RUN go build -o api cmd/api/*

EXPOSE 8000

CMD ["/app/api"]