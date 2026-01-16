FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder

ARG TARGETOS TARGETARCH

RUN apk add --no-cache git ca-certificates gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${TARGETARCH}
RUN go build -o todo-app

FROM alpine:latest

ENV TZ=Asia/Shanghai

RUN apk add --no-cache tzdata ca-certificates && \
    ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone

WORKDIR /app

COPY --from=builder /app/todo-app .
COPY --from=builder /app/config /app/config

RUN addgroup -g 1000 appgroup && \
    adduser -D -u 1000 -G appgroup appuser && \
    chown -R appuser:appgroup /app

USER appuser

EXPOSE 8091

CMD ["./todo-app"]
