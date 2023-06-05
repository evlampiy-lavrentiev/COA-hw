FROM golang:latest

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 2000

ENTRYPOINT ["go", "run", "app/proxy/main.go"]