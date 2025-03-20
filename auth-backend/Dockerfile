FROM --platform=$BUILDPLATFORM golang:alpine AS build
WORKDIR /app
COPY . .
#RUN apk update --no-cache && apk add make
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/dga cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/bin/dga /app/dga
EXPOSE 80
ENTRYPOINT ["/app/dga"]
