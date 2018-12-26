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


