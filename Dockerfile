FROM alpine
RUN adduser -S -D -h /app ketch && \
    apk update && \
    apk --no-cache upgrade && \
    apk --no-cache add ca-certificates && \
    mkdir -p /app/bin
USER ketch
COPY ./bin/linux-amd64/ketch /app/bin
WORKDIR /app
ENTRYPOINT ["/app/bin/ketch"]
