# Bind-mount geth socket to /var/run/geth.ipc
FROM perigord
LABEL maintainer="Maxwell Koo <maxk@polyswarm.io>"

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
    mv perigord-docker.yaml perigord.yaml && \
    perigord build && \
    go get . && \
    go install

EXPOSE 31337

ENTRYPOINT $GOPATH/bin/polyswarm
