FROM golang:1.22.2

WORKDIR /usr/src/app

COPY . .
RUN go build -v /usr/src/app/cmd/main.go

CMD ["./main"]
