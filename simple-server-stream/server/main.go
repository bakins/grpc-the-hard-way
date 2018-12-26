package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	pb "github.com/bakins/grpc-the-hard-way/routeguide/pb/routeguide"
	"github.com/bakins/grpc-the-hard-way/simple-unary-stream/stream"
)

func main() {
	address := flag.String("address", "127.0.0.1:10000", "address to listen for HTTP requests")

	flag.Parse()

	mux := http.NewServeMux()

	// path is <package>.<service>/<method>
	mux.HandleFunc("/routeguide.RouteGuide/ListFeatures", handleListFeatures)

	s := http.Server{
		Addr:    *address,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

var savedFeatures = []pb.Feature{
	{
		Name: "Patriots Path, Mendham, NJ 07945, USA",
		Location: &pb.Point{
			Latitude:  407838351,
			Longitude: -746143763,
		},
	},
	{
		Name: "101 New Jersey 10, Whippany, NJ 07981, USA",
		Location: &pb.Point{
			Latitude:  408122808,
			Longitude: -743999179,
		},
	},
}

func handleListFeatures(w http.ResponseWriter, r *http.Request) {
	var req pb.Rectangle

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

	for _, f := range savedFeatures {

		if err := writer.Write(&f); err != nil {
			http.Error(w, "failed to write response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
