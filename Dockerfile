FROM golang:1.24rc2-alpine3.20 AS builder


COPY . /github.com/quitedevil/chatservice/sourse/
WORKDIR /github.com/quitedevil/chatservice/sourse/

RUN go mod download
RUN go build -o ./bin/service_linux cmd/chatservice/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/quitedevil/chatservice/sourse/bin/chat_service .

CMD [ "./chat_service" ]