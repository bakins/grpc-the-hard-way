# gRPC client side streaming

This is an example of a gRPC service that uses client side streaming.

The gRPC message framing has been abstracted into a [message](./message) library. These examples support [gRPC payload compression](https://github.com/grpc/grpc/blob/d8662f5704ec6f03122943f9baa5ed07b88a1fdf/doc/compression.md).

The comments in the code should hopefully be detailed enough to give you a rough idea about what is happening.

## Testing

To test, open two terminal windows and cd into the root of your clone of this repository.


### client
To test the client, run the following in the first terminal window:

```shell
go run ./standard/client-stream-server/main.go
```

and in the second:

```shell
go run ./client-stream/client/main.go
```

You should see something like:

```
2018/12/28 11:32:19 response: hello 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
2018/12/28 11:32:19 grpc-status: 0
2018/12/28 11:32:19 grpc-message:
```

### server

To test the server, run the following in the first terminal window:

```shell
go run ./client-stream/server/main.go
```

and in the second:

```shell
go run ./standard/client-stream-client/main.go -gzip
```

You should see something like:

```
2018/12/28 11:33:00 Greeting: hello 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
```


