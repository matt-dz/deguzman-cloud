FROM --platform=$BUILDPLATFORM golang:alpine AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY . .
RUN apk update && apk add make
RUN make

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/bin/kube-deploy /app/kube-deploy
EXPOSE 80
ENTRYPOINT ["/app/kube-deploy"]
