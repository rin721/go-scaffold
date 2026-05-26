# syntax=docker/dockerfile:1

FROM golang:1.24-bookworm AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/go-scaffold-server ./cmd/server

FROM debian:bookworm-slim AS runtime

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates curl tzdata \
    && rm -rf /var/lib/apt/lists/* \
    && groupadd --system --gid 10001 app \
    && useradd --system --uid 10001 --gid app --home-dir /app --shell /usr/sbin/nologin app

WORKDIR /app

COPY --from=build /out/go-scaffold-server /app/go-scaffold-server
COPY configs/config.example.yaml /app/configs/config.example.yaml
COPY deploy/config.production.example.yaml /app/configs/config.yaml
COPY configs/locales /app/configs/locales

RUN mkdir -p /app/data /app/logs \
    && chown -R app:app /app

USER app

EXPOSE 9999

ENV REI_CONFIG_PATH=/app/configs/config.yaml

ENTRYPOINT ["/app/go-scaffold-server"]
CMD ["server", "--config=/app/configs/config.yaml"]
