FROM golang:1.19beta1 as builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux go build -a -o playground main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/playground .
USER 65532:65532

ENTRYPOINT ["/playground"]