FROM golang:1.19-alpine as builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bookmarkey cmd/server/bookmarkey/main.go

FROM scratch
COPY --from=builder /build/bookmarkey .

ENTRYPOINT [ "./bookmarkey" ]
CMD ["serve", "--http=0.0.0.0:8080"]