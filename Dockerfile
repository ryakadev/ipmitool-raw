FROM golang:1.24.3-alpine

RUN apk add --no-cache ipmitool

WORKDIR /app
COPY . .

RUN go build -o ipmi-api

EXPOSE 8080

CMD ["./ipmi-api"]