FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add tzdata ca-certificates

# 使用 buildx 时自动注入 TARGETARCH 变量
ARG TARGETARCH

COPY build/spage-linux-${TARGETARCH} /app/server

EXPOSE 8888

RUN chmod +x /app/server

ENTRYPOINT ["/app/server"]