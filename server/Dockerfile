FROM golang:latest as builder
WORKDIR /workspace

COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /workspace/server .

CMD ["./server"]