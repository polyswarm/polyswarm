#!/bin/bash

DEV_DIR=$HOME/etherdev

mkdir -p $DEV_DIR/etc $DEV_DIR/keystore
cp ./genesis.json $DEV_DIR
cp ./keystore/* $DEV_DIR/keystore/
echo -n blah > $DEV_DIR/etc/pw

geth --datadir $DEV_DIR --nodiscover --maxpeers 0 --mine --minerthreads 1 --rpc init $DEV_DIR/genesis.json
geth --datadir $DEV_DIR --nodiscover --maxpeers 0 --mine --minerthreads 1 --rpc --rpcaddr "0.0.0.0" --rpcapi "eth,web3,personal,net" console

geth --exec "personal.unlockAccount(eth.accounts[0], \"blah\", 0)" attach http://localhost:8545
geth --exec "personal.unlockAccount(eth.accounts[1], \"blah\", 0)" attach http://localhost:8545
geth --exec "personal.unlockAccount(eth.accounts[2], \"blah\", 0)" attach http://localhost:8545
