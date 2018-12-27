package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/net/http2"

	"github.com/bakins/grpc-the-hard-way/server-stream/message"
	pb "github.com/bakins/grpc-the-hard-way/services/greetings"
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

	req := pb.GreetingRequest{
		Name: name,
	}

	var buf bytes.Buffer
	if err := message.Write(&buf, &req, true); err != nil {
		log.Fatalf("failed to prepare request: %v", err)
	}

	// path is /<package>.<service>/<method>
	r, err := http.NewRequest(http.MethodPost,
		"http://"+*address+"/greetings.Greeter/ShareGreetings",
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

	// for a server side streaming reply, we loop until we get EOF
READ:
	for {
		var response pb.GreetingReply

		err := message.Read(resp.Body, &response)

		switch err {
		case nil:
			// keep going
		case io.EOF:
			// done reading
			break READ
		default:
			log.Fatalf("failed to read response: %v", err)
		}

		log.Printf("response: %s", response.GetMessage())

	}

	status := 0
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
