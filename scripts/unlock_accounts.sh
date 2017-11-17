#!/bin/bash

geth --exec "personal.unlockAccount(eth.accounts[0], \"blah\", 0)" attach http://localhost:8545
geth --exec "personal.unlockAccount(eth.accounts[1], \"blah\", 0)" attach http://localhost:8545
geth --exec "personal.unlockAccount(eth.accounts[2], \"blah\", 0)" attach http://localhost:8545
