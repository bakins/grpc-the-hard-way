package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	pb "github.com/bakins/grpc-the-hard-way/services/greetings"
	"github.com/bakins/grpc-the-hard-way/unary-v2/message"
)

func main() {
	address := flag.String("address", "127.0.0.1:50051", "address to listen for HTTP requests")

	flag.Parse()

	mux := http.NewServeMux()

	// path is /<package>.<service>/<method>
	mux.HandleFunc("/greetings.Greeter/CrowdGreeting", handleCrowdGreeting)

	s := http.Server{
		Addr:    *address,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

func handleCrowdGreeting(w http.ResponseWriter, r *http.Request) {
	useGzip := r.Header.Get("Grpc-Encoding") == "gzip"

	// in this example, we report any error as a gRPC error
	// in the grpc-status trailer

	// Set correct content type
	w.Header().Set("Content-Type", "application/grpc+proto")

	// We must include trailers for status and message.
	w.Header().Set("Trailer", "grpc-status, grpc-message")
	w.Header().Set("grpc-accept-encoding", "gzip")

	if useGzip {
		w.Header().Set("Grpc-Encoding", "gzip")
	}

	w.WriteHeader(200)

	var names []string
READ:
	for {
		var req pb.GreetingRequest

		err := message.Read(r.Body, &req)

		switch err {
		case nil:
			// keep reading
		case io.EOF:
			// done
			break READ
		default:
			log.Printf("failed to read from stream: %v", err)
			// https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
			// https://github.com/grpc/grpc-go/blob/master/codes/codes.go
			// 10 => Aborted
			w.Header().Set("Grpc-Status", strconv.Itoa(10))
			w.Header().Set("Grpc-Message", err.Error())
			return
		}
		names = append(names, req.GetName())
	}

	resp := pb.GreetingReply{
		Message: "hello " + strings.Join(names, ", "),
	}

	if err := message.Write(w, &resp, useGzip); err != nil {
		http.Error(w, "failed to write response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Grpc-Status", strconv.Itoa(0))
	w.Header().Set("Grpc-Message", "ok")
}
