ARG ALPINE_VERSION=3.10
FROM golang:1.12.6-alpine3.10

RUN apk add git
ENV GO111MODULE=auto

 #override:
ENV LOOKUP_SERVER localhost
ENV LOGSTASH_HOST logstash
ENV LOGSTASH_PORT 5044

WORKDIR /go/src/app

COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT [ "app -n google.com" ]

# CMD ["-n",${LOOKUP_SERVER}]