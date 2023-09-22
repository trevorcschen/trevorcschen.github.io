FROM golang:1.21.1-alpine AS builder

WORKDIR /go/src/app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o server ./cmd/app

FROM alpine
LABEL language="golang"
LABEL org.opencontainers.image.source https://github.com/trevorcschen/go-ms-example

WORKDIR /app
COPY --from=builder /go/src/app/ /app/

# CMD [ "/app/server" ]
ENTRYPOINT [ "./server" ]