# build stage
FROM golang:1.16 AS builder
LABEL maintainer="ivan.rodkevich"

ENV GO111MODULE=auto
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /go/src/github.com/rodkevich/gmp-tickets
COPY ../.. .
RUN go get ./...
WORKDIR /go/src/github.com/rodkevich/...

RUN go build -ldflags="-s -w" ./server.go

# server stage
FROM scratch
LABEL maintainer="ivan.rodkevich/app"

COPY --from=builder /go/src/github.com/rodkevich/... .

EXPOSE 12300

CMD ["./server"]
