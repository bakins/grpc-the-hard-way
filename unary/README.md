# simple gRPC unary implementation

This contains naive, straightforward implementations of a gRPC client and server.  The service implemented is [helloworld](../services/helloworld) from the [gRPC examples](https://github.com/grpc/grpc-go/tree/7b141362910abb44ee44416797a8da21659d5ae4/examples/helloworld/helloworld).

[protoc](https://github.com/golang/protobuf) was used to generate the protobuf handling code.  The gRPC generators are not used.

These examples do not support [gRPC payload compression](https://github.com/grpc/grpc/blob/d8662f5704ec6f03122943f9baa5ed07b88a1fdf/doc/compression.md).

* [client](./client) implements a basic gRPC service client using the [stdlib http client](https://golang.org/pkg/net/http/) and a few helpers.
* [server](./server) implements a basic gRPC server using the [stdlib http server](https://golang.org/pkg/net/http/) and a few helpers.

The comments in the code should hopefully be detailed enough to give you a rough idea about what is happening.

## Testing

To test, open two terminal windows and cd into the root of your clone of this repository.

### client
To test the client, run the following in the first terminal window:

```shell
go run ./standard/greeter_server/main.go
```

and in the second:

```shell
go run ./unary/client/main.go
```

You should see something like:

```
2018/12/26 14:50:03 grpc-status: 0
2018/12/26 14:50:03 grpc-message:
2018/12/26 14:50:03 response: Hello world
```

### server

To test the server, run the following in the first terminal window:

```shell
go run ./unary/server/main.go
```

and in the second:

```shell
go run ./standard/greeter_client/main.go
```

You should see something like:

```
2018/12/26 14:51:27 Greeting: hello world
```


