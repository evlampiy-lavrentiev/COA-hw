FROM golang:latest

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV MULTICAST_ADDRESS 239.0.0.1:54321

ENTRYPOINT ["go", "run", "app/worker/main.go"]
