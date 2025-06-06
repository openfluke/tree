FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY . ./
RUN touch .env

# 👇 Compile all Go files in the current directory
RUN go build -o openfluke .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/openfluke /app/openfluke
COPY ./templates ./templates
COPY ./static ./static
COPY .env .env
COPY cert.pem cert.pem
COPY key.pem key.pem

EXPOSE 8080
EXPOSE 4443

ENTRYPOINT ["/app/openfluke"]
