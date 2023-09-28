FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -ldflags="-s -w" -o /app/washbonus ./cmd/washbonus

FROM alpine

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app

COPY environment/firebase /app/firebase

COPY internal/migrations /app/internal/migrations
COPY --from=builder /app/washbonus /app/washbonus

EXPOSE 8080
CMD ["/app/washbonus"]
