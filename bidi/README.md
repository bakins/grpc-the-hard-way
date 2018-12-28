# gRPC client side streaming

This is an example of a bidirectional gRPC service.

The gRPC message framing has been abstracted into a [message](./message) library. These examples support [gRPC payload compression](https://github.com/grpc/grpc/blob/d8662f5704ec6f03122943f9baa5ed07b88a1fdf/doc/compression.md).

The comments in the code should hopefully be detailed enough to give you a rough idea about what is happening.

## Testing

To test, open two terminal windows and cd into the root of your clone of this repository.


### client
To test the client, run the following in the first terminal window:

```shell
go run ./standard/bidi-server/main.go
```

and in the second:

```shell
go run ./bidi/client/main.go
```

You should see something like:

```
2018/12/28 13:19:54 response: مرحبا 0
2018/12/28 13:19:54 response: こんにちは 1
2018/12/28 13:19:54 response: Hello 2
2018/12/28 13:19:54 response: γεια σας 3
2018/12/28 13:19:54 response: Bonjour 4
2018/12/28 13:19:54 response: مرحبا 5
2018/12/28 13:19:54 response: こんにちは 6
2018/12/28 13:19:54 response: γεια σας 7
2018/12/28 13:19:54 response: Hello 8
2018/12/28 13:19:54 response: Hello 9
2018/12/28 13:19:54 grpc-status: 0
2018/12/28 13:19:54 grpc-message:
```

### server

To test the server, run the following in the first terminal window:

```shell
go run ./bidi/server/main.go
```

and in the second:

```shell
go run ./standard/bidi-client/main.go -gzip
```

You should see something like:

```
2018/12/28 13:09:41 Greeting: مرحبا 0
2018/12/28 13:09:41 Greeting: こんにちは 1
2018/12/28 13:09:41 Greeting: Hello 2
2018/12/28 13:09:41 Greeting: γεια σας 3
2018/12/28 13:09:41 Greeting: Bonjour 4
2018/12/28 13:09:41 Greeting: مرحبا 5
2018/12/28 13:09:41 Greeting: こんにちは 6
2018/12/28 13:09:41 Greeting: γεια σας 7
2018/12/28 13:09:41 Greeting: Hello 8
2018/12/28 13:09:41 Greeting: Hello 9
```


