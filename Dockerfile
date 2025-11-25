# Build stage
FROM golang:1.21-alpine AS build

WORKDIR /app
COPY main.go .
RUN go build -o server main.go

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/server .
EXPOSE 1234
CMD ["./server"]
