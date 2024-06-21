FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go
EXPOSE 3000
CMD ["./app"]
