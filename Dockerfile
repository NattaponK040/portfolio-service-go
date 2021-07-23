FROM golang:1.16.5 as build
COPY . /app
WORKDIR /app
RUN go build -o /app .
