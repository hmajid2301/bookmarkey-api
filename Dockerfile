FROM golang:1.19-alpine as builder

WORKDIR /build
RUN apk update && apk upgrade && \
	apk add --no-cache ca-certificates && \
	update-ca-certificates

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bookmarkey cmd/server/bookmarkey/main.go

FROM scratch
COPY --from=builder /build/bookmarkey .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./bookmarkey" ]
CMD ["serve", "--http=0.0.0.0:8080", "--encryptionEnv=PB_ENCRYPTION_KEY"]