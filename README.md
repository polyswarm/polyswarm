# PolySwarm Alpha

*Early Stages of development*

## Introduction

This repository contains a Perigord project which represents the end-to-end
(early) test network implementation of Bounties as described in our whitepaper
(https://polyswarm.io/polyswarm-whitepaper.pdf). Offers are coming soon. All
contracts in this repo are subject to change.

## Bounties

Bounties in PolySwarm give enterprises the ability to submit artifacts to
PolySwarm and receive responses from experts on the maliciousness of artifacts.
Enterprises can leverage the code in `bounty/poster.go` as a Golang based
interface for posting suspect artifacts to PolySwarm. Experts can use the same
interface to stream bounties and post their assertions (opinions) on files of
interest for which they have expertise.

### Posting (Enterprise/Users)

There are several components of each bounty posted:

* Bounty fee which experts must pay to record their assertions
* Bounty amount which is the initial reward offered to experts for examining the
  file
* Artifact information: hash, URL to fetch
* Deadline
* Bounty GUID

Enterprises can leverage the public methods in `bounty/poster.go` as a Golang
based interface for posting suspect artifacts to PolySwarm.

### Assertions

Experts can use the same interface to stream bounties and post their assertions
(opinions) on files of interest for which they have expertise.

An event-based API for streaming bounties is provided.

Assertions against posted bounties consist of:

* A bid against the bounty
* A determination of malicious-or-not (boolean)
* Optional metadata (such as e.g. malware family) which provides value-add to
  the bounty poster

## Offers

Offers are a work in progress, but will represent a direct offer from an
enterprise to a security expert to analyze an artifact. To issue an offer, the
enterprise will open a Raiden-style channel with the expert and issue zero or
more offers.

## Running

### Native

Steps:

0) Install Perigord with all dependencies (https://github.com/polyswarm/perigord )
1) Launch geth testnet with `./scripts/launch_geth_testnet.sh`
2) Build bindings with `perigord build`
3) Run tests with `go test`
4) Run web interface with `go run main.go`

### Docker

Steps:

0) Build docker image with included Dockerfile (`docker/Dockerfile`)
1) Launch geth testnet with `./scripts/launch_geth_testnet.sh`
2) Run docker image, mounting geth IPC socket to `/var/run/geth.ipc` in the
container
3) Run tests with `go test`
4) Run web interface with `go run main.go`
