FROM golang:1.23.1 as builder

WORKDIR /server

COPY . .

RUN go build -o serve .

EXPOSE 8080

CMD ["./serve"]
