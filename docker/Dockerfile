FROM alpine:3.5

# app workdir
WORKDIR /go/src/avalon
ADD ./binary/avalon /go/src/avalon/
ADD ./binary/start.sh /go/src/avalon/

RUN set -ex && apk add --update bash ca-certificates \
            && rm -rf /var/cache/apk/*

# runtime configs
EXPOSE 8080
ENTRYPOINT ["./start.sh"]
CMD ["run"]