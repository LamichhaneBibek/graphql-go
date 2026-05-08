FROM golang:1.26-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air"]