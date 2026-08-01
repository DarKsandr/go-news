FROM ubuntu:24.04

RUN apt-get update \
    && apt-get install -y --no-install-recommends golang-go ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && go install github.com/swaggo/swag/cmd/swag@latest

ENV PATH="/root/go/bin:${PATH}"

WORKDIR /app
