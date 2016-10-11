FROM golang:1.7

ADD . /go/src/github.com/chrisbodhi/zonked

# RUN apk update && apk upgrade && \
#    apk add --no-cache bash git openssh

RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

ENTRYPOINT revel run github.com/chrisbodhi/zonked dev 8080

EXPOSE 8080