FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go build -o Image

FROM alpine

WORKDIR /app

ENV PORT="4000"

COPY --from=builder /build/Image /app/Image

CMD ["./Image"]