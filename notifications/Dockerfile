FROM golang:1.15.13-alpine as builder

WORKDIR /app

COPY . .

# RUN apk add gcc musl-dev
RUN apk --no-cache add ca-certificates
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags -ldflags="-w -s" .

FROM busybox

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/notifications .
COPY --from=builder /app/config/vendors.toml ./config/
COPY --from=builder /app/.env .

EXPOSE 8085

CMD [ "./notifications" ]