FROM golang:1.18
WORKDIR /go/src/go-products-api
COPY . .
RUN make build
CMD ["./bin/server"]