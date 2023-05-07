FROM golang:1.19-bullseye as builder

WORKDIR /app
ADD ./ ./
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM debian:10-slim
RUN apt update && apt install -y ca-certificates
COPY --from=builder /app/main /app/main
EXPOSE 8080

ENTRYPOINT [ "/app/main" ]