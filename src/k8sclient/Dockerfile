FROM golang:latest as builder

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY pkg/ ./pkg/
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly  -v -o /tmp/k8sclient

FROM alpine:3 as runner
RUN apk add --no-cache ca-certificates
COPY --from=builder /tmp/k8sclient /go/bin/k8sclient

ENV LOG_LEVEL='info'
ENV SERVER_PORT=8080
EXPOSE ${SERVER_PORT}
ENTRYPOINT ["/go/bin/k8sclient"]
