FROM --platform=$BUILDPLATFORM golang:alpine AS build
WORKDIR /app
COPY . .
RUN apk update && apk add make
RUN make

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/bin/dss /app/dss
EXPOSE 80
ENTRYPOINT ["/app/dss"]
