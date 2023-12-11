FROM mysql:latest



FROM golang:latest
WORKDIR /app
COPY . .
RUN go get -u github.com/go-sql-driver/mysql
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]