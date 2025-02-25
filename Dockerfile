FROM golang:1.24rc2-alpine3.20 AS builder

COPY . /chatservice/sourse
WORKDIR /chatservice/sourse

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service_chat cmd/chatserver/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder chatservice/sourse/service_chat .
ADD .env .
ADD keys/ .
RUN chmod +x service_chat
CMD [ "./service_chat" ]