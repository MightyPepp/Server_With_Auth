FROM golang:1.26.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go get github.com/ilyakaznacheev/cleanenv@v1.5.0 && go build -o server ./cmd/api/main.go

EXPOSE 8443

CMD ["./server"]