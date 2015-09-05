FROM golang:1.4-cross
RUN apt-get update && apt-get install -y --no-install-recommends openssh-client
RUN go get github.com/mitchellh/gox
RUN go get github.com/aktau/github-release
RUN go get -u github.com/golang/glog
RUN go get -u github.com/constabulary/gb/...

ENV GOPATH /go/
WORKDIR /go/src/github.com/nlamirault/abraracoursix

ADD . /go/
ADD vendor /go/
