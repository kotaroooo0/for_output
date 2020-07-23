FROM golang:1.14-alpine as builder

WORKDIR /go/src/github.com/kotaroooo0/for_output/depends_on

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main

FROM alpine:latest

COPY --from=builder /go/src/github.com/kotaroooo0/for_output/depends_on/main /main

ENTRYPOINT ["/main"]
