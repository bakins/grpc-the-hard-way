# simple server stream implementation 
This contains simple implementations of a gRPC service with server side streaming

The gRPC message framing has been abstracted into a [message](./message) library. These examples support [gRPC payload compression](https://github.com/grpc/grpc/blob/d8662f5704ec6f03122943f9baa5ed07b88a1fdf/doc/compression.md).

The comments in the code should hopefully be detailed enough to give you a rough idea about what is happening.

## Testing

To test, open two terminal windows and cd into the root of your clone of this repository.

### client
To test the client, run the following in the first terminal window:

```shell
go run ./standard/server-stream-server/main.go
```

and in the second:

```shell
go run ./server-stream/client/main.go
```

You should see something like:

```
2018/12/27 10:18:34 response: Hola world
2018/12/27 10:18:35 response: こんにちは world
2018/12/27 10:18:36 response: γεια σας world
2018/12/27 10:18:37 response: Hallo world
2018/12/27 10:18:38 response: مرحبا world
2018/12/27 10:18:39 response: Hello world
2018/12/27 10:18:40 response: Bonjour world
2018/12/27 10:18:41 grpc-status: 0
2018/12/27 10:18:41 grpc-message:
```

### server

To test the server, run the following in the first terminal window:

```shell
go run ./server-stream/server/main.go
```

and in the second:

```shell
go run ./standard/server-stream-client/main.go -gzip
```

You should see something like:

```
2018/12/27 10:19:18 Greeting: Hello world
2018/12/27 10:19:18 Greeting: Hallo world
2018/12/27 10:19:18 Greeting: γεια σας world
2018/12/27 10:19:18 Greeting: Hola world
2018/12/27 10:19:19 Greeting: Bonjour world
2018/12/27 10:19:19 Greeting: مرحبا world
2018/12/27 10:19:19 Greeting: こんにちは world
```


