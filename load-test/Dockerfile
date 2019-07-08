ARG ALPINE_VERSION=3.10
FROM golang:1.12.6-alpine3.10

RUN apk add git

WORKDIR /go/src/load-test

ENV GO111MODULE=on

COPY . . 

RUN go mod init load-test

RUN go get -d -v ./...

RUN go build

 #override:
ENV LOOKUP_SERVER localhost
ENV LOGSTASH_HOST logstash
ENV LOGSTASH_PORT 5044

ENTRYPOINT [ "./load-test" ]

CMD ["-n",${LOOKUP_SERVER}]