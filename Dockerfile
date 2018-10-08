FROM alpine:latest

COPY main /usr/local/bin/

RUN chmod +x /usr/local/bin/main

ENTRYPOINT [ "main" ]
