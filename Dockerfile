FROM --platform=$BUILDPLATFORM golang:1.18-alpine as BUILD

WORKDIR /re

RUN apk update && \
  apk --no-cache add make git

# Copy go.mod and go.sum first and download for caching go modules
COPY go.mod go.sum ./

RUN go mod download

# Copy the files from host
COPY . .

ARG TARGETARCH TARGETOS
ENV GOOS=${TARGETOS} GOARCH=${TARGETARCH} LEDGER_ENABLED=false BUILD_TAGS=muslc
RUN make build

FROM alpine:latest

ENV USER_HOME /re

RUN addgroup --gid 1025 reuser && adduser --uid 1025 -S -G reuser reuser -h "$USER_HOME"

USER reuser

WORKDIR $USER_HOME

COPY --from=BUILD /re/bin/red /usr/bin/red