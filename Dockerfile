FROM golang:1.20-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -C cmd/server/bookmarkey .

FROM alpine:3.17.3
COPY --from=builder /build/cmd/server/bookmarkey/bookmarkey ./bookmarkey
COPY --from=flyio/litefs:0.3 /usr/local/bin/litefs /usr/local/bin/litefs

ADD etc/litefs.yml /etc/litefs.yml

RUN apk add fuse

ENTRYPOINT litefs mount -- ./bookmarkey serve --http=0.0.0.0:8080 --encryptionEnv=PB_ENCRYPTION_KEY
