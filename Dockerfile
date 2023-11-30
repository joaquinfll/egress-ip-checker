FROM registry.access.redhat.com/ubi9/go-toolset:1.20 AS builder

USER 10001

COPY . /app/src

WORKDIR /app/src
RUN go build -o /app/src/server

FROM registry.access.redhat.com/ubi9/micro:9.3

USER 10001

WORKDIR /app
COPY --from=builder /app/src/server /app/server

ENTRYPOINT ["/app/server"]