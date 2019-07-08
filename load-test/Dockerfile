ARG ALPINE_VERSION=3.10
FROM golang:1.12.6-alpine3.10

ENV GO111MODULE=on

 #override:
ENV LOOKUP_SERVER localhost
ENV LOGSTASH_HOST logstash
ENV LOGSTASH_PORT 5044

WORKDIR /go/src/app

COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT [ "load-test" ]

CMD ["-n",${LOOKUP_SERVER}]