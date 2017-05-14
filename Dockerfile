FROM alpine
RUN apk add --update ca-certificates && rm -rf /var/cache/apk/* /tmp/*
EXPOSE 3000
COPY cepinator /
ENTRYPOINT ./cepinator
