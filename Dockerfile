FROM alpine:latest

WORKDIR /app

COPY goshirase /app/goshirase

ENTRYPOINT ["/app/goshirase"]