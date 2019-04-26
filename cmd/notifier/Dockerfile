FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD notifier /app
ENTRYPOINT ["./notifier"]