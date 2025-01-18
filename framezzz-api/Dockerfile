FROM golang:1.23 AS build

WORKDIR /app

COPY framezzz-api/go.mod ./
COPY framezzz-api/go.sum ./

RUN go mod download

COPY framezzz-api/cmd/server/main.go ./cmd/server/main.go
COPY framezzz-api/cmd/server/docs/ ./cmd/server/docs/
COPY framezzz-api/pkg/ ./pkg/
COPY framezzz-api/internal/ ./internal/
COPY framezzz-api/config/ ./config/

RUN CGO_ENABLED=0 go build -o ./server ./cmd/server/main.go

# Stage 2
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 ./app/server ./
COPY framezzz-api/migrations/ ./migrations/
COPY .env .env
COPY framezzz-api/logs/ ./logs/

EXPOSE 7312

CMD ["./server"]