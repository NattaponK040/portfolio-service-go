FROM golang:1.16.5
COPY . /app
WORKDIR /app
RUN go build -o /app .
