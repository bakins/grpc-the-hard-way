package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/http2"

	pb "github.com/bakins/grpc-the-hard-way/services/helloworld/helloworld"
)

const (
	defaultName = "world"
)

func main() {
	address := flag.String("address", "127.0.0.1:50051", "address of server")
	flag.Parse()

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// configure transport to allow HTTP over plain text connections
	t := &http2.Transport{
		AllowHTTP: true,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		},
	}

	client := http.Client{
		Transport: t,
	}

	// create and marshal our request
	req := pb.HelloRequest{
		Name: name,
	}

	body, err := proto.Marshal(&req)
	if err != nil {
		log.Fatalf("failed to marshal request: %v", err)
	}

	// gRPC requests and responses include a prefix before each message
	// The prefix is
	//  - one byte flag that denotes if the the message is compressed or not
	//  - four bytes that are an unsigned 32 bit integer. This is the length
	//    of the message.
	// The message follows this prefix.
	// See https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md

	// create a prefix message.
	// this version does not support compression, so flag is always 0.
	// using this syntax rather than make just to be more explicit.
	prefix := []byte{0, 0, 0, 0, 0}
	binary.BigEndian.PutUint32(prefix[1:], uint32(len(body)))

	// create our request body
	// 5 is length of prefix
	request := make([]byte, 5+len(body))
	copy(request, prefix)
	copy(request[5:], body)

	// path is <package>.<service>/<method>
	resp, err := client.Post(
		"http://"+*address+"/helloworld.Greeter/SayHello",
		"application/grpc+proto",
		bytes.NewBuffer(request),
	)

	if err != nil {
		log.Fatalf("failed to POST request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("unexpected status code: %d %s", resp.StatusCode, resp.Status)
	}

	// read the response, starting with the prefix
	prefix = []byte{0, 0, 0, 0, 0}
	_, err = resp.Body.Read(prefix)
	if err != nil {
		log.Fatalf("failed to read prefix: %v", err)
	}

	// determine the length of the message.  Future versions
	// should ensure this is a valid length - ie, not 0 and not greater
	// than a configured maximum size
	length := binary.BigEndian.Uint32(prefix[1:])

	body = make([]byte, length)
	_, err = resp.Body.Read(body)
	if err != nil {
		log.Fatalf("failed to read body: %v", err)
	}

	var helloResponse pb.HelloReply
	if err := proto.Unmarshal(body, &helloResponse); err != nil {
		if err != nil {
			log.Fatalf("failed to unmarshal body: %v", err)
		}
	}

	log.Printf("response: %s", helloResponse.GetMessage())

	// must read until EOF to ensure trailers are read.
	// there should be no data left before trailers.
	if _, err = resp.Body.Read([]byte{}); err != io.EOF {
		log.Fatalf("unexpected error: %v", err)
	}

	status := 0
	// this is set in a trailer sent by the server.
	grpcStatus := resp.Trailer.Get("Grpc-Status")
	if grpcStatus != "" {
		s, err := strconv.Atoi(grpcStatus)
		if err != nil {
			log.Fatalf("failed to parse grpc-status %s: %v", grpcStatus, err)
		}
		status = s
	}

	log.Printf("grpc-status: %d", status)
	// Note: grpc-message may not be sent if status is 0/ok
	log.Printf("grpc-message: %s", resp.Trailer.Get("Grpc-Message"))
	if status != 0 {
		log.Fatal("unexpected grpc status")
	}
}
