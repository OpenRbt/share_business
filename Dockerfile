FROM golang:1-buster as builder

ARG GITHUB_USER=$GITHUB_USER
ARG GITHUB_PASS=$GITHUB_PASS

RUN printf "machine github.com\n\tlogin %s\n\tpassword %s" ${GITHUB_USER} ${GITHUB_PASS} >> ~/.netrc && \
    apt-get update -yqq && \
    apt-get install -yqq make git build-essential && \
    apt-get clean && \
    apt-get autoremove -yqq && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

WORKDIR /app

COPY cmd/wash_bonus/.env .
COPY ./ ./

RUN go build -o ./bin/wash-bonus ./cmd/main/*


FROM alpine:3  

COPY --from=builder /app/bin/wash-bonus /bin/wash-bonus
COPY --from=builder /app/migration/*.sql /migration/

RUN apk add --no-cache libc6-compat ca-certificates

CMD ["/bin/wash-bonus"] 
