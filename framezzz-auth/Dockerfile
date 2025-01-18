FROM golang:1.23 AS build

WORKDIR /app

COPY framezzz-auth/go.mod ./
COPY framezzz-auth/go.sum ./

RUN go mod download

COPY framezzz-auth/models/ ./models/
COPY framezzz-auth/controllers/ ./controllers/
COPY framezzz-auth/routes/ ./routes/
COPY framezzz-auth/utils/ ./utils/
COPY framezzz-auth/middlewares/ ./middlewares/
COPY framezzz-auth/main.go ./main.go

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o ./server ./main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/server /root/server

COPY .env .env

EXPOSE 7313

CMD ["./server"]
