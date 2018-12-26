package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/http2"

	pb "github.com/bakins/grpc-the-hard-way/routeguide/pb/routeguide"
	"github.com/bakins/grpc-the-hard-way/simple-unary-stream/stream"
)

func main() {
	address := flag.String("address", "127.0.0.1:10000", "address of server")
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

	var buf bytes.Buffer

	s := stream.New(&buf, true)

	req := pb.Rectangle{
		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
	}

	if err := s.Write(&req); err != nil {
		log.Fatalf("failed to prepare request: %v", err)
	}

	// path is <package>.<service>/<method>
	r, err := http.NewRequest(http.MethodPost,
		"http://"+*address+"/routeguide.RouteGuide/ListFeatures",
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

	reader := stream.New(resp.Body, resp.Header.Get("Grpc-Encoding") == "gzip")

	for {
		var feature pb.Feature
		err := reader.Read(&feature)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to ListFeatures: %v", err)
		}
		log.Println(feature)
	}

	// must read the entire body before you can read trailers
	log.Println(resp.Trailer)
}
