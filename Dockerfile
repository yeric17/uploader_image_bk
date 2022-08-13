
#Build stage
FROM golang:1.17.2-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/public ./public
RUN apk update && apk add bash
EXPOSE 4000
CMD [ "/app/main" ]

