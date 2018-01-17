FROM alpine:latest

RUN mkdir /scripts

COPY main /scripts/

RUN chmod +x /scripts/main

CMD /scripts/main
