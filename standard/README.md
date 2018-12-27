# standard

This directory contains clients and servers  written using the ["standard" gRPC Go implementation](https://godoc.org/google.golang.org/grpc).

* [greeter_server](./greeter_server) is a modified version of and [example](https://github.com/grpc/grpc-go/tree/7b141362910abb44ee44416797a8da21659d5ae4/examples/helloworld/greeter_server) from https://github.com/grpc/grpc-go - this version supports message compression
* [greeter_client](./greater_client) is a modified version of and [example](https://github.com/grpc/grpc-go/tree/7b141362910abb44ee44416797a8da21659d5ae4/examples/helloworld/greeter_client) from https://github.com/grpc/grpc-go - this version supports message compression via a flag.
* [server-stream-server](./server-stream-server) simple server side streaming version of hello world.
* [server-stream-client](./server-stream-client) client of the streaming helloworld service.

This contains code licensed under the Apache 2 license, copyrighted by the gRPC authors.