FROM golang:1.15

WORKDIR /app/bin/

COPY time-server .

CMD ["/app/bin/time-server"]
