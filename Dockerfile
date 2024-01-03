FROM golang:1.21.5-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

EXPOSE 8080

WORKDIR /dist

COPY . .

RUN go get -v -t -d ./...
RUN go build -o torvalds cmd/cli/main.go

FROM scratch

COPY --from=builder /dist/torvalds /

CMD ["/torvalds"]