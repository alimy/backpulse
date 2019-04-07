FROM golang:alpine AS binaryBuilder
# Install build deps
RUN apk --no-cache --no-progress add --virtual build-deps build-base git

# Build project
WORKDIR /go/src/github.com/backpulse/core
COPY . .
RUN make build

FROM alpine:latest
# Install system utils & runtime dependencies
ADD https://github.com/tianon/gosu/releases/download/1.11/gosu-amd64 /usr/sbin/gosu
RUN chmod +x /usr/sbin/gosu \
  && echo http://dl-2.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories \
  && apk --no-cache --no-progress add bash s6 shadow

# Configure LibC Name Service
COPY hack/docker/nsswitch.conf /etc/nsswitch.conf

# Copy target app from binaryBuilder stage
WORKDIR /app/backpulse
COPY hack/docker docker
COPY --from=binaryBuilder /go/src/github.com/backpulse/core/backpulse .

# Finalize s6 configure
RUN ./docker/finalize.sh

# Configure Docker Container
ENV MONGODB_URI mongodb://mongodb:27017
ENV DATABASE backpulse

VOLUME ["/data"]

# 8000: backend data interface agent.
# 3000: dashboard ui agent.
# 3001: frontpulse ui agent.
EXPOSE 8000
EXPOSE 3000
EXPOSE 3001

ENTRYPOINT ["/app/backpulse/docker/start.sh"]
CMD ["/bin/s6-svscan", "/app/backpulse/docker/s6/"]