FROM golang:1.13-alpine AS builder
RUN apk add --no-cache bash gcc make musl-dev
WORKDIR /build
COPY . .
RUN make

FROM alpine:latest
RUN apk add --no-cache bash nano
COPY --from=builder /build/bin/cmd/gdcrm /usr/local/bin/
EXPOSE 4441 4441/udp 4441 4441/tcp 4449 4449/tcp
COPY ./docker-entrypoint.sh /
RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]
