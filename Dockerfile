FROM registry.access.redhat.com/ubi9/go-toolset:1.20 as builder

CMD mkdir -p /app/src
COPY . /app/src

CMD go build -o /app/src/server

FROM registry.access.redhat.com/ubi9/micro:1.20

CMD mkdir -p /app
COPY --from=builder /app/src/server /app/server

ENTRYPOINT ["/app/server"]