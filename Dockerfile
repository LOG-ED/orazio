FROM golang:1.9.1
WORKDIR /go/src/github.com/log-ed/orazio
RUN go get github.com/tools/godep  
ADD . /go/src/github.com/log-ed/orazio
RUN godep restore
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/log-ed/orazio/app .
COPY --from=0 /go/src/github.com/log-ed/orazio/tmpl/satirae.html tmpl/satirae.html
CMD ["./app"]