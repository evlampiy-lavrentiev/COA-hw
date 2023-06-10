FROM golang:latest

WORKDIR /app
COPY go.mod go.sum ./

RUN apt update --yes && apt install --yes protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go


COPY . .
RUN protoc --go_out=. app/worker/types/anek.proto
RUN go mod download



ENTRYPOINT ["go", "run", "app/worker/main.go"]
