FROM alpine
RUN apk add --update ca-certificates && rm -rf /var/cache/apk/* /tmp/*
EXPOSE 3000
COPY ./dist/cepinator_Linux_x86_64/cepinator /
ENTRYPOINT ./cepinator
