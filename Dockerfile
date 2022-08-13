
#Build stage
FROM golang:1.17.2-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# #Run stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main ./bin/sh
COPY --from=builder /app/.env .
EXPOSE 4000
CMD [ "/app/main" ]

