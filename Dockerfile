FROM golang:1.22.1-alpine AS builder

WORKDIR /dist

COPY . .

RUN go get -v -t -d ./...
RUN go build -o torvalds cmd/action/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /dist/torvalds /

CMD ["/torvalds"]