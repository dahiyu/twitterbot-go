# Build image
FROM golang:latest AS builder

# Install dependencies
WORKDIR /go/src/twitter-bot
RUN go get github.com/ChimeraCoder/anaconda && \
    go get github.com/joho/godotenv && \
    go get github.com/mmcdole/gofeed && \
    go get github.com/jasonlvhit/gocron

# Build modules
COPY main.go .
COPY .env .
RUN GOOS=linux CGO_ENABLED=0 go build main.go

#--

# Production image
FROM busybox
WORKDIR /opt/twitter-bot/bin

# Deploy modules
COPY --from=builder /go/src/twitter-bot .
ENV TZ=Asia/Tokyo
COPY crontab /var/spool/cron/crontabs/root
# ENTRYPOINT ["/opt/twitter-bot/bin/main"]
# CMD ["crond" "-f", "-d", "8"]
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
CMD crond -f -d 8