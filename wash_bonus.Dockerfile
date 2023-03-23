FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY cmd/wash_bonus/go.mod .
COPY cmd/wash_bonus/go.sum .

RUN go mod download

COPY cmd/wash_bonus/ .

RUN go build -ldflags="-s -w" -o /app/wash_bonus .

FROM alpine

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app

COPY environment/certs/ /app/certs/
COPY environment/firebase /app/firebase

COPY cmd/wash_bonus/migrations /app/migrations
COPY --from=builder /app/wash_bonus /app/wash_bonus

EXPOSE 8080
CMD ["/app/wash_bonus"]
