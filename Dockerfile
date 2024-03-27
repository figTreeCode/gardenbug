FROM golang:alpine as build 

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd/app

FROM alpine:latest

WORKDIR /app

RUN apk add ca-certificates

COPY --from=build /app/main /app/main 

EXPOSE 8080

CMD ["./main"]

