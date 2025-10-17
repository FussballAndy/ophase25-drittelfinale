FROM golang:1.25.2-alpine3.21

WORKDIR /app

COPY . .

RUN go get

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]