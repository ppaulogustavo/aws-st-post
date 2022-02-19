FROM golang:1.17

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go mod download
RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]