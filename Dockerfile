FROM golang:bullseye AS builder
WORKDIR /app
COPY src/ .
RUN go mod download
RUN go build -o ./server ./main.go


FROM golang:bullseye AS runner
WORKDIR /app

COPY --from=builder /app/server /app/server
COPY --from=builder /app/config/sms.yml /app/config/sms.yml
RUN ln -snf /usr/share/zoneinfo/Asia/Tehran /etc/localtime && echo Asia/Tehran > /etc/timezone
EXPOSE 8000
ENTRYPOINT ["/app/server"]  