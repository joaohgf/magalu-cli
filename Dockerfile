FROM golang:1.26 AS builder

WORKDIR /src

COPY go.mod go.sum ./
COPY vendor ./vendor
COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o /out/tarefeiro ./cmd/tarefeiro

FROM alpine:3.22

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /out/tarefeiro /usr/local/bin/tarefeiro

ENTRYPOINT ["tarefeiro"]
CMD ["--help"]

