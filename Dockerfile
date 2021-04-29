FROM alpine:latest

LABEL app="hfctl"
LABEL maintainer="manigandan.jeff@gmail.com"
LABEL version="v1.0.0"
LABEL description="HelloFresh Recipe Stats Calculator CLI"

RUN mkdir -p /app && apk update && apk add --no-cache ca-certificates
RUN mkdir -p /app/test
WORKDIR /app
# This require the project to be built first before copying,
# else docker build will fail
COPY hfctl /app/
COPY test/json_file.log /app/test/

VOLUME /app

ENTRYPOINT ["/app/hfctl"]
# CMD "/app/hfctl"
