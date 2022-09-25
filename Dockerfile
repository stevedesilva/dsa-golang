FROM golang:alpine
RUN mkdir /app
COPY commonsense/server /app
COPY go.mod /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]