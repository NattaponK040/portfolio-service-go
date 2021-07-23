FROM golang:1.16

WORKDIR /go-portfolio-service
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-portfolio-service"]