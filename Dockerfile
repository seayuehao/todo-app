# ---------- build stage ----------
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o todo-app

FROM alpine:latest

ENV TZ=Asia/Shanghai

RUN apk add --no-cache tzdata ca-certificates && \
    ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone

WORKDIR /app

COPY --from=builder /app/todo-app .

RUN addgroup -g 1000 appgroup && \
    adduser -D -u 1000 -G appgroup appuser && \
    chown -R appuser:appgroup /app

USER appuser

EXPOSE 8091

CMD ["./todo-app"]
