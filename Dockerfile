FROM golang:1.15.6 as build
COPY . /app
WORKDIR /app
RUN go build -o /app/app.go