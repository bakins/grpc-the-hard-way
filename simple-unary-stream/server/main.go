package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	pb "github.com/bakins/grpc-the-hard-way/helloworld/pb/helloworld"
	"github.com/bakins/grpc-the-hard-way/simple-unary-stream/stream"
)

func main() {
	address := flag.String("address", "127.0.0.1:50051", "address to listen for HTTP requests")

	flag.Parse()

	mux := http.NewServeMux()

	// path is <package>.<service>/<method>
	mux.HandleFunc("/helloworld.Greeter/SayHello", handleSayHello)

	s := http.Server{
		Addr:    *address,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

func handleSayHello(w http.ResponseWriter, r *http.Request) {
	var req pb.HelloRequest

	useGzip := r.Header.Get("Grpc-Encoding") == "gzip"
	reader := stream.New(r.Body, useGzip)

	if err := reader.Read(&req); err != nil {
		http.Error(w, "failed to read request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Set correct content type
	w.Header().Set("Content-Type", "application/grpc+proto")

	// We must include trailers for status and message.
	w.Header().Set("Trailer", "grpc-status, grpc-message")
	w.Header().Set("grpc-status", strconv.Itoa(0))
	w.Header().Set("grpc-message", "ok")
	w.Header().Set("grpc-accept-encoding", "gzip")

	if useGzip {
		w.Header().Set("Grpc-Encoding", "gzip")
	}

	writer := stream.New(w, useGzip)

	resp := pb.HelloReply{
		Message: "hello " + req.GetName(),
	}

	if err := writer.Write(&resp); err != nil {
		http.Error(w, "failed to write response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
