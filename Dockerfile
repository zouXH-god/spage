# build
FROM reg.liteyuki.icu/dockerhub/golang:1.24.2-alpine3.21 AS builder

ENV TZ=Asia/Chongqing

WORKDIR /app

RUN apk --no-cache add build-base git tzdata

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main  \
    -ldflags="-X 'github.com/LiteyukiStudio/spage/config.CommitHash=$(git rev-parse HEAD)'  \
    -X 'github.com/LiteyukiStudio/spage/config.BuildTime=$(date "+%Y-%m-%d %H:%M:%S")'"  \
    main.go

# production
FROM reg.liteyuki.icu/dockerhub/alpine:latest AS prod

ENV TZ=Asia/Chongqing

WORKDIR /app

RUN apk --no-cache add tzdata ca-certificates libc6-compat libgcc libstdc++

COPY --from=builder /app/main /app/main

EXPOSE 8888

RUN chmod +x ./main

ENTRYPOINT ["./main"]