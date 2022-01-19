FROM golang:1.17.5-alpine AS builder
RUN apk update && apk --no-cache add ca-certificates tzdata \
  bash git openssh gcc g++ pkgconfig build-base curl \
  zlib-dev librdkafka-dev pkgconf && rm -rf /var/cache/apk/*
# Copy app and run go mod.
WORKDIR /go/src/github.com/maurodanieldev/quasar-oper-fire
COPY ./go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -tags musl -a -installsuffix cgo -o app .
# Copy and run the app.
FROM alpine:3.11.11
WORKDIR /usr/
COPY --from=builder /go/src/github.com/maurodanieldev/quasar-oper-fire/app .
CMD /usr/app