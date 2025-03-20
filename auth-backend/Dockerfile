FROM --platform=$BUILDPLATFORM golang:alpine AS build
WORKDIR /app
COPY . .
RUN apk update && apk add make
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH make

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/bin/dga /app/dga
EXPOSE 80
ENTRYPOINT ["/app/dga"]
