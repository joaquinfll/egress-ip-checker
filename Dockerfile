FROM registry.access.redhat.com/ubi9/go-toolset:1.20 AS builder


RUN ls -larth /
RUN ls -larth .
COPY . .

RUN go mod tidy && \
    go build -o server main.go

FROM registry.access.redhat.com/ubi9/ubi-micro:9.3

USER 1000

WORKDIR /app
COPY --from=builder /opt/app-root/src/server /app/server

ENTRYPOINT ["/app/server"]