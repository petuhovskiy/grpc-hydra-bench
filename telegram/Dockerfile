FROM    golang:alpine AS builder
RUN     apk update && apk add bash ca-certificates git gcc g++ libc-dev binutils file
WORKDIR /usr/src/app
COPY    go.mod .
COPY    go.sum .
RUN     go mod download
COPY    . .
RUN     go build -o /app .


FROM alpine:latest
RUN  apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /app ./
CMD  ["./app"]