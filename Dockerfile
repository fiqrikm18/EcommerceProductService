FROM golang:1.24-alpine as builder
WORKDIR /ecommerce
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /ecommerce/product-service ./cmd/app

FROM alpine:edge as production
WORKDIR /app
COPY --from=builder /ecommerce/product-service .
COPY docs docs
COPY .env .env
RUN apk --no-cache add ca-certificates tzdata
ENTRYPOINT ["/app/product-service"]
