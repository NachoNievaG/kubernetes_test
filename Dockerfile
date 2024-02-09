# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /main
EXPOSE 8080

# Run
CMD ["/main"]
