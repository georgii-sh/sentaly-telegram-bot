FROM golang:1.15 as builder

LABEL maintainer="Georgii Sh <georgii@sentaly.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Run unit/integrations testing before build
RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/config.yml .

EXPOSE 8080

CMD ["./main"] 