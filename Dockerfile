FROM golang:1.24rc2-alpine3.20 AS builder

COPY . /chatservice/sourse
WORKDIR /chatservice/sourse

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service_chat cmd/chatserver/main.go

FROM alpine:latest

WORKDIR /root/app
COPY --from=builder chatservice/sourse/service_chat .
COPY --from=builder chatservice/sourse/.env .
WORKDIR /root/app/keys
COPY --from=builder chatservice/sourse/keys .

WORKDIR /root/app
RUN chmod +x service_chat
CMD [ "./service_chat" ]