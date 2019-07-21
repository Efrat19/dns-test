FROM go_base

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

CMD ["-n",${LOOKUP_SERVER}]
ENTRYPOINT [ "app" ]
