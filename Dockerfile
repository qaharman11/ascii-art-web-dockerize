FROM golang:latest

LABEL maintainer="qaharman11 and Nurlybek316728"

WORKDIR /app

COPY . .

RUN cd cmd; go build -o main

EXPOSE 8080

CMD ["./cmd/main"]