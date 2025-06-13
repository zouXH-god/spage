FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add tzdata ca-certificates

ARG TARGETARCH

COPY build/spage-linux-${TARGETARCH} /app/server

RUN chmod +x /app/server

ENTRYPOINT ["/app/server"]