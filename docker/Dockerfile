FROM golang:1.12.9-alpine AS builder

WORKDIR /go/src/github.com/dynastymasra/avalon

# install dependecies
RUN set -ex \
    && apk add --update bash git curl \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && git config --global http.https://gopkg.in.followRedirects true \
    && rm -rf /var/cache/apk/*
COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure -v --vendor-only

## build linux app source code
COPY . ./
RUN GOOS=linux go build -tags=main -o avalon

FROM alpine:3.10.1
RUN set -ex && apk add --update bash ca-certificates tzdata \
    && rm -rf /var/cache/apk/*

# app
WORKDIR /app
COPY --from=builder /go/src/github.com/dynastymasra/avalon/avalon /app/

## runtime configs
EXPOSE 8080
ENTRYPOINT ["./avalon"]
CMD ["start"]
