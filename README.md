# gRPC The Hard Way

gRPC clients and servers in Go using stdlib and a few helpers. Avoid https://godoc.org/google.golang.org/grpc

## Motivation

gRPC in Go is often criticized for being bloated and/or too complicated.
However, gRPC the protocol is relatively simply.  I wanted to see
how far I could get using stdlib to implement gRPC in Go.

Originally, I was going to write a reusable, alternative gRPC Go package,
but I decided that would be:
* more work than I wanted to do in a simple side project
* be filled with all of my opinions

So, I decided to do straight-forward examples of simple gRPC services in Go.

## Status

The code here is, hopefully, fairly straight-forward.  It is not "pretty" by any means. I have not attempted to make good abstractions. I've tried
to be fairly "brute force" on purpose.

The clients and servers I have written are confirmed to work with clients and
servers written using the ["standard" gRPC Go implementation](https://godoc.org/google.golang.org/grpc).

## Overview

The [gRPC Protocol document](https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md) provides the details of the gRPC protocol.  A summary is that gRPC is protocol buffer encoded message over HTTP2 - and the messages are prefixed with a small amont of metadata.

## Contents

Each directory listed should have a README.md with more information.

* [standard-grpc](./standard-grpc) contains clients and servers written using the ["standard" gRPC Go implementation](https://godoc.org/google.golang.org/grpc). These are used to test my implementations.
* [simple-unary](./simple-unary) contains naive implementations of a simple gRPC service.
* [simple-unary-stream](./simple-unary-stream) abstracts the message handling into a common library.


The clients and servers use h2c (HTTP2 over plaintext connections). They do not support TLS.  In production, something like [envoy](https://www.envoyproxy.io) handles hop-to-hop encryption for me.

## Usage

These have been tested with Go 1.11.4.  Clone this into your Go path like:

```bash
mkdir -p $(go env GOPATH)/src/github.com/bakins
cd $(go env GOPATH)/src/github.com/bakins
git clone https://github.com/bakins/grpc-the-hard-way.git grpc-the-hard-way
cd grpc-the-hard-way
```

Any commands ran in any docs assume they are being ran from the root of the
repository clone.

## TODO

* streaming server
* streaming client
* bi-directional communication

