FROM golang
COPY . /app
WORKDIR /app
RUN go build -o /app .
