FROM golang:1.21.5-alpine AS builder

WORKDIR /dist

COPY . .

RUN go get -v -t -d ./...
RUN go build -o torvalds cmd/cli/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /dist/torvalds /

CMD ["/torvalds"]