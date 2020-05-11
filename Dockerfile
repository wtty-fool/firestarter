FROM golang:1.14 AS builder

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN mkdir -p /opt/firestarter
WORKDIR /opt/firestarter

COPY main.go .
COPY go.mod .
COPY go.sum .

RUN go get -v -t -d ./...
RUN go build -v .

# ------------------------------------
FROM alpine:3.11

COPY --from=builder /opt/firestarter/firestarter /bin/firestarter

ENTRYPOINT ["/bin/firestarter"]


