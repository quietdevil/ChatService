FROM alpine:3.21.3

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.24.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD migrations/*.sql migrations/
ADD migrator.sh .
ADD .env .

RUN chmod +x migrator.sh

ENTRYPOINT ["bash", "migrator.sh"]