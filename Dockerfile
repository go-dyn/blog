FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0  \
    GOARCH="amd64" \
    GOOS=linux

WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build --ldflags "-extldflags -static" -o hero .

FROM alpine
#FROM gcr.io/distroless/static-debian12
WORKDIR /zero
COPY --from=builder /build/hero /zero/hero
COPY --from=builder /build/database/ /zero/database/
COPY --from=builder /build/public/ /zero/public/
COPY --from=builder /build/storage/ /zero/storage/
COPY --from=builder /build/.env /zero/.env

ENTRYPOINT ["/zero/hero"]