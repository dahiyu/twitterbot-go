# Build image
FROM golang:latest AS builder
# Install dependencies
WORKDIR /go/src/twitter-bot
RUN go get github.com/ChimeraCoder/anaconda && \
    go get github.com/mmcdole/gofeed
# Build modules
COPY main.go .
COPY .env .
RUN GOOS=linux CGO_ENABLED=0 go build main.go

# --

# ca-certificates
FROM alpine AS certificates
RUN apk update && apk add ca-certificates

# --

# Production image
FROM busybox
WORKDIR /opt/twitter-bot/bin
# certificates
COPY --from=certificates /etc/ssl/certs /etc/ssl/certs
# Deploy modules
COPY --from=builder /go/src/twitter-bot .
ENV TZ=Asia/Tokyo
COPY crontab /var/spool/cron/crontabs/root
# CMD ["crond" "-f", "-d", "8"]
CMD crond -f -d 8
