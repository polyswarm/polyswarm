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

RUN set -x && \
    cd /tmp && \
    git clone https://github.com/raiden-network/raiden.git --branch v0.2.0 && \
    cd raiden && \
    pip install --upgrade -r requirements.txt && \
    python setup.py develop

RUN mkdir -p /go/src/github.com/polyswarm/polyswarm
ENV GOPATH=/go
ADD . /go/src/github.com/polyswarm/polyswarm

ADD /scripts/keystore /keystore

WORKDIR /go/src/github.com/polyswarm/polyswarm
RUN go get .

EXPOSE 8080

CMD ["go", "run", "main.go"]
