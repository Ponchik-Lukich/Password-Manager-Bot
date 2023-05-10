FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main_local main_local.go

EXPOSE 3000

CMD ["./main_local"]
