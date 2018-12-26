package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/net/http2"

	pb "github.com/bakins/grpc-the-hard-way/helloworld/pb/helloworld"
	"github.com/bakins/grpc-the-hard-way/simple-unary-stream/stream"
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

	var buf bytes.Buffer

	s := stream.New(&buf, true)

	req := pb.HelloRequest{
		Name: name,
	}

	if err := s.Write(&req); err != nil {
		log.Fatalf("failed to prepare request: %v", err)
	}

	// path is <package>.<service>/<method>
	r, err := http.NewRequest(http.MethodPost,
		"http://"+*address+"/helloworld.Greeter/SayHello",
		bytes.NewBuffer(buf.Bytes()),
	)

	if err != nil {
		log.Fatalf("failed to create request: %v", err)
	}

	r.Header.Set("Content-Type", "application/grpc+proto")
	r.Header.Set("grpc-encoding", "gzip")

	resp, err := client.Do(r)

	if err != nil {
		log.Fatalf("failed to POST request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("unexpected status code: %d %s", resp.StatusCode, resp.Status)
	}

	s = stream.New(resp.Body, resp.Header.Get("Grpc-Encoding") == "gzip")

	var helloResponse pb.HelloReply

	if err := s.Read(&helloResponse); err != nil {
		log.Fatalf("failed to read response: %v", err)
	}

	log.Printf("response: %s", helloResponse.GetMessage())

	status := 0
	grpcStatus := resp.Trailer.Get("Grpc-Status")
	if grpcStatus != "" {
		s, err := strconv.Atoi(grpcStatus)
		if err != nil {
			log.Fatalf("failed to parse grpc-status %s: %v", grpcStatus, err)
		}
		status = s
	}

	// must read until EOF to ensure trailers are read
	if _, err = ioutil.ReadAll(resp.Body); err != nil && err != io.EOF {
		log.Fatalf("unexpected error: %v", err)
	}

	log.Printf("grpc-status: %d", status)
	// Note: grpc-message may not be sent if status is 0/ok
	log.Printf("grpc-message: %s", resp.Trailer.Get("Grpc-Message"))
	if status != 0 {
		log.Fatal("unexpected grpc status")
	}

}
