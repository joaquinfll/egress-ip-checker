FROM registry.access.redhat.com/ubi9/go-toolset:1.20 AS build

CMD mkdir -p /app/src
COPY . /app/src

WORKDIR /app/src
RUN go mod download
RUN go build -o server ./main.go

FROM registry.access.redhat.com/ubi9-micro:9.3

CMD mkdir -p /app
COPY --from=build /app/src/server /app/server

ENTRYPOINT ["/app/server"]