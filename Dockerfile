# mostly stolen from:
# https://github.com/ardanlabs/service/blob/a5fe29b9132c04eef4f6b14d10e12cdba1ab9a0a/dockerfile.sales-api

# build the go binary
FROM golang:1.13-alpine as go_builder
ENV CGO_ENABLED 0

WORKDIR /code

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN go build -mod=readonly -ldflags '-w -extldflags "-static"' .

# main app container
FROM graphiteapp/docker-graphite-statsd:latest
ARG BUILD_DATE
ARG VCS_REF

# inject the go binary into the container as a runit service
# see: https://sanjeevan.co.uk/blog/running-services-inside-a-container-using-runit-and-alpine-linux/
RUN mkdir -p /etc/service/bts-stats
COPY --from=go_builder /code/bts-trading-stats /etc/service/bts-stats/run

ENTRYPOINT ["/entrypoint"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="bts-trading-stats" \
      org.opencontainers.image.authors="Samo Ratnik <samo.ratnik@gmail.com>" \
      org.opencontainers.image.source="https://github.com/samotarnik/bts-trading-stats" \
      org.opencontainers.image.revision="${VCS_REF}"
