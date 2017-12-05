# Bind-mount geth socket to /var/run/geth.ipc
FROM perigord
LABEL maintainer="Maxwell Koo <mjkoo90@gmail.com>"

RUN set -x && \
    apt-get update && \
    apt-get -y install \
        libssl-dev \
        python3 \
        python3-pip && \
    rm -rf /var/lib/apt/lists/*

RUN set -x && \
    cd /usr/src && \
    git clone https://github.com/raiden-network/microraiden.git && \
    cd microraiden && \
    pip3 install -e microraiden

RUN mkdir -p /go/src/github.com/polyswarm/polyswarm
ENV GOPATH=/go
ADD . /go/src/github.com/polyswarm/polyswarm

ADD /scripts/keystore /keystore

WORKDIR /go/src/github.com/polyswarm/polyswarm
RUN go get .

EXPOSE 8080

CMD ["go", "run", "main.go"]
