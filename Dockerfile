FROM golang:1.24rc2-alpine3.20 AS builder


COPY . /chatservice/sourse
WORKDIR /chatservice/sourse

RUN go mod download
RUN go build cmd/chatservice/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/quitedevil/chatservice/sourse/service_linux .

CMD [ "./chat_service" ]