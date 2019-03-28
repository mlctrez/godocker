FROM golang:1.12 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY cli cli
COPY pkg pkg

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cli/godocker/main.go



FROM scratch

COPY --from=builder /build/app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY static static

CMD ["./app"]
