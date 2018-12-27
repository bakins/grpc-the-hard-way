# unary implementation with message handling library

This contains simple implementations of a gRPC service.  This is a modification of the [unary](../unary) examples.

The gRPC message framing has been abstracted into a [message](./message) library. These examples support [gRPC payload compression](https://github.com/grpc/grpc/blob/d8662f5704ec6f03122943f9baa5ed07b88a1fdf/doc/compression.md).

The comments in the code should hopefully be detailed enough to give you a rough idea about what is happening.

## Testing

To test, open two terminal windows and cd into the root of your clone of this repository.

These examples are similar to the [unary](../unary) examples, but they use a library that implements payload compression.

### client
To test the client, run the following in the first terminal window:

```shell
go run ./standard/greeter_server/main.go
```

and in the second:

```shell
go run ./unary-v2/client/main.go
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
go run ./unary-v2/server/main.go
```

and in the second:

```shell
go run ./standard/greeter_client/main.go -gzip
```

You should see something like:

```
2018/12/26 14:51:27 Greeting: hello world
```


