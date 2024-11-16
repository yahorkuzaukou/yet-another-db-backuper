FROM golang:1.23.3-alpine

COPY --chown=app . .

RUN go build

CMD ["./app"]