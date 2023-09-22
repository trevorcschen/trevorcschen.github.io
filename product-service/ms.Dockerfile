FROM golang:1.21.1-alpine
RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o server ./cmd/app

CMD [ "/app/server" ]

