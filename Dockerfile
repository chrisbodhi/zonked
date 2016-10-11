# FROM golang:1.7-alpine

# ADD . /go/src/zonked

# RUN go install github.com/chrisbodhi/zonked

# ENTRYPOINT /go/bin/zonked
FROM golang:onbuild
EXPOSE 8080