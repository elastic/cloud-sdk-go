FROM golang:1.18-stretch

RUN GO111MODULES=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
