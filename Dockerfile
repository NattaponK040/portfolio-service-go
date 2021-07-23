FROM golang:1.12.17
RUN go version

ADD . /go/src/app
WORKDIR /go/src/app

# Expose 8080
# Gin will use the PORT env var
ENV PORT 8080
EXPOSE 8080

# Compile app
RUN go build -o main .
# Run app
CMD ["/go/src/app/main"]