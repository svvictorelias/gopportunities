FROM golang:1.22.2-bookworm as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

ARG PORT
ENV PORT=$PORT
EXPOSE $PORT
CMD ["./main"]