FROM golang:1.16.12 AS build
ENV CGO_ENABLED=0
# Build
WORKDIR /app
COPY . .
RUN go build -mod vendor -o bin/crud-server

FROM alpine:3.9
# Run
WORKDIR /
ENV GODEBUG madvdontneed=1
EXPOSE 80
COPY --from=build /app/bin/crud-server /app/service

ENTRYPOINT ["/app/service"]