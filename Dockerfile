FROM golang:1.12.6-alpine3.10

RUN apk add git

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY . . 

RUN go get -d -v ./...

RUN go install -v ./...


 #override:
ENV LOOKUP_SERVER localhost

CMD ["dns-test","-n",${LOOKUP_SERVER}]