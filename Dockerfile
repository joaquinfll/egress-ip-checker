FROM registry.access.redhat.com/ubi9/go-toolset:1.20 AS builder

USER 1000

ARG GOPATH=/workspace/output/go
ENV GOPATH=${GOPATH}

WORKDIR /workspace/output/src
RUN go mod tidy && go build -o server main.go

FROM registry.access.redhat.com/ubi9/ubi-micro:9.3

USER 1000

WORKDIR /app
COPY --from=builder /workspace/output/src/server /app/server

ENTRYPOINT ["/app/server"]