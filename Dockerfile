FROM golang:1.14
WORKDIR /opt/sandbox-docker-go
COPY ./src .
RUN go build -o app main.go
CMD ["/opt/sandbox-docker-go/app"]