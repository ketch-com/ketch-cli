FROM alpine
ARG TARGETOS TARGETARCH
RUN adduser -S -D -h /ketch ketch && \
    apk update && \
    apk --no-cache upgrade && \
    apk --no-cache add ca-certificates
USER ketch
COPY .build/$TARGETOS-$TARGETARCH/ketch /usr/local/bin
COPY docker/ketch/docker-entrypoint.sh /usr/local/bin
WORKDIR /ketch
ENTRYPOINT ["docker-entrypoint.sh"]
