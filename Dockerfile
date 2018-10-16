FROM golang:1.11.1

MAINTAINER tMatSuZ

ADD main.go /go/src/github.com/tMatSuZ/goshirase/

RUN go get github.com/dghubble/go-twitter/twitter
RUN go get github.com/dghubble/oauth1
RUN go get github.com/coreos/pkg/flagutil
RUN go get github.com/jinzhu/gorm
RUN go install github.com/tMatSuZ/goshirase

ENTRYPOINT /go/bin/goshirase