FROM golang:1.13 as builder
WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
# Copy the go source
COPY ./*.go ./
COPY default.key .
COPY default.pem .
# Build
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o server .

FROM alpine
RUN apk update && apk add --no-cache curl
WORKDIR /workspace
# Collect binaries and assets
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=builder /workspace/server .
COPY --from=builder /workspace/default.key .
COPY --from=builder /workspace/default.pem .
CMD /workspace/server
