FROM golang:1.18-alpine AS builder
WORKDIR /data
COPY go.mod go.sum ./
COPY cmd ./cmd
COPY pkg ./pkg
RUN go mod download
RUN ls -l ./
RUN go build -o hmip_sync ./cmd/*

FROM alpine:3  
RUN apk --no-cache add ca-certificates
RUN addgroup -g 1000 hmip_sync && adduser -DH -h / -u 1000 -G hmip_sync hmip_sync
COPY --from=builder /data/hmip_sync /
RUN chown hmip_sync:hmip_sync /hmip_sync
USER hmip_sync
CMD ["/hmip_sync"]  
