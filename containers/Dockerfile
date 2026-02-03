# =========================
# BUILD STAGE
# =========================
FROM golang:1.25.6 AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build explícito apontando para main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o api-booking .

# =========================
# RUNTIME STAGE
# =========================
FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

ENV TZ=America/Sao_Paulo

COPY --from=builder /app/api-booking .

EXPOSE 8000

CMD ["./api-booking"]
