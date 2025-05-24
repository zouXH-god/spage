# build
FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /app

RUN apk --no-cache add build-base git tzdata

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o server \
    -ldflags="-X './config.CommitHash=$(git rev-parse HEAD)'  \
    -X './config.BuildTime=$(date "+%Y-%m-%d %H:%M:%S")'"  \
    ./cmd/server

# production
FROM alpine:latest AS prod

WORKDIR /app

RUN apk --no-cache add tzdata ca-certificates libc6-compat libgcc libstdc++

COPY --from=builder /app/server /app/server

EXPOSE 8888

RUN chmod +x ./server

ENTRYPOINT ["./server"]