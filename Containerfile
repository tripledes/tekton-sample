FROM golang:1.16
WORKDIR /go/src/github.com/tripledes/web-quotes/
COPY app.go go.mod go.sum ./
COPY pkg ./pkg
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/tripledes/web-quotes/app ./
RUN chmod 555 ./app
EXPOSE 8080
CMD ["./app"]  