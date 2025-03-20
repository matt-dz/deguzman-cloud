FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go build -o bin/dga cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/bin/dga /app/dga
EXPOSE 80
ENTRYPOINT ["/app/dga"]
