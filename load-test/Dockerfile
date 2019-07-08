ARG ALPINE_VERSION=3.10
FROM golang:1.12.6-alpine3.10

RUN apk add git

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY . . 

RUN go mod init ${PWD}

RUN go build

 #override:
ENV LOOKUP_SERVER localhost
ENV LOGSTASH_HOST logstash
ENV LOGSTASH_PORT 5044

ENTRYPOINT [ "./load-test" ]

CMD ["-n",${LOOKUP_SERVER}]