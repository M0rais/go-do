FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /app/

WORKDIR /app/cmd/server

RUN go build -o /app/main .

EXPOSE 8080

CMD ["/app/main"]
