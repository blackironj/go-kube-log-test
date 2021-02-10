FROM library/golang:1.15-alpine AS builder

ENV GO111MODULE=on

ENV APP_DIR /tmp/go-kube-log-test
WORKDIR $APP_DIR
COPY . .

RUN go mod tidy 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -s' -o /go-kube-log-test .

FROM scratch
COPY --from=builder /go-kube-log-test /go-kube-log-test
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/go-kube-log-test"]
