FROM golang:alpine AS binaryBuilder

RUN apk add --no-cache git

WORKDIR /go/src/github.com/backpulse/core
COPY . .
RUN make build

FROM alpine:latest
WORKDIR /app/backpulse
COPY --from=binaryBuilder /go/src/github.com/backpulse/core/backpulse .

EXPOSE 8000
CMD ["./backpulse", "serve"]