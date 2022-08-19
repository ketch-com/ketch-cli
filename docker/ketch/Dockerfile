FROM alpine
RUN adduser -S -D -h /ketch ketch && \
    apk update && \
    apk --no-cache upgrade && \
    apk --no-cache add ca-certificates
USER ketch
COPY .build/linux-amd64/ketch /usr/local/bin
COPY docker/ketch/docker-entrypoint.sh /usr/local/bin
WORKDIR /ketch
ENTRYPOINT ["docker-entrypoint.sh"]
