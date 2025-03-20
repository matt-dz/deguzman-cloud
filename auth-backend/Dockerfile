FROM --platform=$BUILDPLATFORM golang:alpine AS build
WORKDIR /app
COPY . .
RUN apk update && apk add make
ARG TARGETOS
ARG TARGETARCH
ARG CGO_ENABLED
RUN TARGETOS=$TARGETOS TARGETARCH=$TARGETARCH CGO_ENABLED=$CGO_ENABLED make

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/bin/dga /app/dga
EXPOSE 80
ENTRYPOINT ["/app/dga"]
