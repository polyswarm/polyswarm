# Bind-mount geth socket to /var/run/geth.ipc
FROM perigord
LABEL maintainer="Maxwell Koo <mjkoo90@gmail.com>"

RUN set -x && \
    apt-get update && \
    apt-get -y install \
        libssl-dev \
        python \
        python-pip && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir -p /go/src/github.com/polyswarm/polyswarm
ENV GOPATH=/go
ADD . /go/src/github.com/polyswarm/polyswarm

ADD /scripts/keystore /keystore

WORKDIR /go/src/github.com/polyswarm/polyswarm
RUN set -x && \
    perigord build && \
    go get . && \
    go install

EXPOSE 5001
EXPOSE 8080

ENTRYPOINT $GOPATH/bin/polyswarm
