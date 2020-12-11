#  Stage 1
FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o server .
# Stage 2   demo-app => server
FROM alpine
#RUN adduser -S -D -H -h /app appuser
#USER appuser
COPY --from=builder /build/ /app/
WORKDIR /app
CMD ["./server"]