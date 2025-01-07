FROM golang:alpine as build
WORKDIR /app
COPY . .
RUN apk update && apk add make
RUN make build

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/bin/qr-generator /app/qr-generator
RUN mkdir /srv/qrcode
ENV SRV_DIR=/srv/qrcode
EXPOSE 80
ENTRYPOINT ["/app/qr-generator"]
