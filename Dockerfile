# Syntax for buildkit
# syntax=docker/dockerfile:1

# Builder Stage
FROM golang:1.22.5 AS builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app main.go


# Runtime Stage
FROM alpine:latest
WORKDIR /app

# Install ca-certificates 
RUN apk add --no-cache ca-certificates

ARG PORT
ARG DATABASE_URL
ARG SIGN_ROLE
ARG SIGN_NAME
ARG SING_NIP

# Copy the binary and .env file from the builder stage
COPY --from=builder /app/app .

# Casbin config
COPY --from=builder /app/config /app/config

#Set Environment Variables
ENV PORT=$PORT
ENV DATABASE_URL=$DATABASE_URL
ENV SIGN_ROLE=$SIGN_ROLE
ENV SIGN_NAME=$SIGN_ROLE
ENV SIGN_NIP=$SIGN_NIP

EXPOSE $PORT
CMD ["./app"]
