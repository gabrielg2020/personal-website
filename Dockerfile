# --- Stage 1: Build ---

FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build server
ENV GOOS=linux CGO_ENABLED=0
RUN go build -ldflags="-w -s" -o bin/server

# --- Stage 2: Run server! ---

FROM alpine:3.19

# Create a non-root user and set permissions
RUN adduser -h /home/app -s /bin/ash -D appuser && chown -R appuser:appuser /home/app && chown 700 /home/app

WORKDIR /home/app

# Copy application binary and supporting files
COPY --from=builder /app/bin/server ./
COPY --from=builder /app/content ./content
COPY --from=builder /app/data ./data
COPY --from=builder /app/views ./views
COPY --from=builder /app/static ./static

USER appuser

EXPOSE 8080

ENV GIN_MODE=release

CMD [ "./server" ]

