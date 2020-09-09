FROM alpine:3.9

RUN apk add -u ca-certificates && apk upgrade \
  && rm -rf /var/*/apk/* /var/log/*

ADD ./health-checker-* /health-checker

USER nobody

ENTRYPOINT ["/health-checker"]