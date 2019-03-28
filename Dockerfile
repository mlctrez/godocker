FROM golang:1.12 as builder

WORKDIR /build
COPY cli cli
COPY pkg pkg
COPY go.mod go.sum ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cli/godocker/main.go

FROM scratch
COPY --from=builder /build/app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/pki/tls/certs/ca-bundle.crt
COPY static static
CMD ["./app"]
