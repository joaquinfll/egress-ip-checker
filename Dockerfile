FROM registry.access.redhat.com/ubi9/go-toolset:1.20 as builder

USER 10001

RUN mkdir -p /app/src
COPY . /app/src


RUN go mod download && go build -o /app/src/server

FROM registry.access.redhat.com/ubi9/micro:9.3

USER 10001

RUN mkdir -p /app
COPY --from=builder /app/src/server /app/server

ENTRYPOINT ["/app/server"]