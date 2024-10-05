FROM golang:1.23.1

WORKDIR /app

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

ENV GOFLAGS=-buildvcs=false

COPY . .

EXPOSE 8080

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

RUN go install github.com/a-h/templ/cmd/templ@latest

ENTRYPOINT CompileDaemon -include="*.templ" -exclude="*_templ.go" -build="templ generate && go build -o serve ." -command="./serve"
