package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"

	"golang.org/x/net/http2"

	pb "github.com/bakins/grpc-the-hard-way/services/greetings"
	"github.com/bakins/grpc-the-hard-way/unary-v2/message"
)

func main() {
	address := flag.String("address", "127.0.0.1:50051", "address of server")
	flag.Parse()

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

	pr, pw := io.Pipe()

	// write in sendNames will block until
	// request body is read
	go sendNames(pw)

	// path is /<package>.<service>/<method>
	r, err := http.NewRequest(http.MethodPost,
		"http://"+*address+"/greetings.Greeter/CrowdGreeting",
		pr,
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

	var response pb.GreetingReply

	if err := message.Read(resp.Body, &response); err != nil {
		log.Fatalf("failed to read response: %v", err)
	}

	log.Printf("response: %s", response.GetMessage())

	// must read until EOF to ensure trailers are read.
	// there should be no data left before the trailers.
	if _, err = resp.Body.Read([]byte{}); err != io.EOF {
		log.Fatalf("unexpected error: %v", err)
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

func sendNames(w io.WriteCloser) {
	// send 10 "names"
	for i := 0; i < 10; i++ {
		req := pb.GreetingRequest{
			Name: strconv.Itoa(i),
		}

		if err := message.Write(w, &req, true); err != nil {
			// in a "real" program, we would handle this error, but for now
			// just exit
			log.Fatalf("failed to write request: %v", err)
		}
	}

	// close to signify that we are done sending requests
	if err := w.Close(); err != nil {
		log.Fatalf("failed to close writer: %v", err)
	}
}
