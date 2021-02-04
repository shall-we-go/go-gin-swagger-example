from golang:1.15

WORKDIR /app/
RUN go get -u github.com/swaggo/swag/cmd/swag

CMD ["/bin/sh", "-c", "tail -f /dev/null"]
