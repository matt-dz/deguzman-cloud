FROM --platform=$BUILDPLATFORM golang:alpine AS build
WORKDIR /app
COPY . .
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -o bin/dga cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/bin/dga /app/dga
EXPOSE 80
ENTRYPOINT ["/app/dga"]
