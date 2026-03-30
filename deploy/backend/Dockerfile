FROM golang:1.24.1-bookworm AS builder

WORKDIR /src/backend

COPY backend/go.mod backend/go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY backend ./

ARG SERVICE_ENTRY
ARG BINARY_NAME

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" \
    -o "/out/${BINARY_NAME}" "${SERVICE_ENTRY}"

FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates tzdata gettext-base \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

ARG BINARY_NAME

COPY --from=builder "/out/${BINARY_NAME}" "/usr/local/bin/${BINARY_NAME}"
COPY deploy/backend/docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh

RUN chmod +x /usr/local/bin/docker-entrypoint.sh \
    && mkdir -p /app/etc /app/templates

ENV CONFIG_TEMPLATE_PATH=/app/templates/config.yaml.tpl
ENV CONFIG_PATH=/app/etc/config.yaml

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]
