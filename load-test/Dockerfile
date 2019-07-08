ARG ALPINE_VERSION=3.10
FROM golang:1.12.6-alpine${ALPINE_VERSION}}

RUN apk add git

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY . . 

RUN go get -d -v ./...

RUN go install -v ./...

 #override:
ENV LOOKUP_SERVER localhost
ENV LOGSTASH_HOST logstash
ENV LOGSTASH_PORT 5044

ENTRYPOINT [ "./load-test" ]

CMD ["-n",${LOOKUP_SERVER}]