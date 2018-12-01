FROM golang:1.11


ADD . /go/src/github.com/Sharykhin/fin-tech

WORKDIR /go/src/github.com/Sharykhin/fin-tech


RUN go get github.com/canthefason/go-watcher && \
    go install github.com/canthefason/go-watcher/cmd/watcher && \
    go get -u gopkg.in/alecthomas/gometalinter.v2 && \
    gometalinter.v2 --install && \
    go get -u github.com/pressly/goose/cmd/goose

EXPOSE 8003

CMD watcher -run github.com/Sharykhin/fin-tech/cmd/api -watch github.com/Sharykhin/fin-tech
