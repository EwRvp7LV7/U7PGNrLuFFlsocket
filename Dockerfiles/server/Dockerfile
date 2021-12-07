FROM golang:1.17-alpine AS builder

RUN apk add upx

WORKDIR /go/src/github.com/EwRvp7LV7/nocloud-tunnel-mesh
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" ./cmd/tunnelserver
RUN upx ./tunnelserver

FROM scratch
WORKDIR /
COPY --from=builder  /go/src/github.com/EwRvp7LV7/nocloud-tunnel-mesh/tunnelserver /tunnelserver

EXPOSE 8080
EXPOSE 8000
LABEL org.opencontainers.image.source https://github.com/EwRvp7LV7/nocloud-tunnel-mesh

ENTRYPOINT ["/tunnelserver"]
