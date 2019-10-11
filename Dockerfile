FROM alpine:latest

RUN addgroup -g 2000 appuser && \
    adduser -D -u 1000 -G appuser appuser

COPY hello /usr/local/bin/

RUN chmod +x /usr/local/bin/hello

USER appuser

ENTRYPOINT [ "hello" ]
