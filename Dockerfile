FROM golang:1.19-buster

RUN GO111MODULES=off go install github.com/go-swagger/go-swagger/cmd/swagger@latest
