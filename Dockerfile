FROM golang:1.12-stretch

RUN GO111MODULES=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
