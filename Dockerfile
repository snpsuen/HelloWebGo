FROM golang:latest
RUN mkdir /app
COPY ./ /app/
WORKDIR /app
RUN go mod init app
RUN go mod tidy
RUN go build -o main .
CMD ["/app/main"]
