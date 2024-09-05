FROM golang:1.21.1 as builder

ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o main .

EXPOSE 8080
CMD ["./main"]
