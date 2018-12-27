package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	mux.HandleFunc("/greetings.Greeter/ShareGreetings", handleShareGreetings)

	s := http.Server{
		Addr:    *address,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

var greetings = []string{
	"Hello",
	"Hola",
	"Hallo",
	"Bonjour",
	"こんにちは",
	"مرحبا",
	"γεια σας",
}

func handleShareGreetings(w http.ResponseWriter, r *http.Request) {
	useGzip := r.Header.Get("Grpc-Encoding") == "gzip"

	var req pb.GreetingRequest

	if err := message.Read(r.Body, &req); err != nil {
		http.Error(w, "failed to read request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Set correct content type
	w.Header().Set("Content-Type", "application/grpc+proto")

	// We must include trailers for status and message.
	w.Header().Set("Trailer", "grpc-status, grpc-message")
	w.Header().Set("grpc-accept-encoding", "gzip")

	if useGzip {
		w.Header().Set("Grpc-Encoding", "gzip")
	}

	w.WriteHeader(200)

	// randomize the list
	tmp := make([]string, len(greetings))
	copy(tmp, greetings)
	rand.Shuffle(len(tmp), func(i, j int) {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	})

	for _, g := range tmp {
		resp := pb.GreetingReply{
			Message: g + " " + req.GetName(),
		}

		if err := message.Write(w, &resp, useGzip); err != nil {
			http.Error(w, "failed to write response: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// in this example, we are flushing on each greeting
		// and sleeping to simulate doing some server side work
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		time.Sleep(time.Millisecond * 250)
	}

	w.Header().Set("Grpc-Status", strconv.Itoa(0))
	w.Header().Set("Grpc-Message", "ok")
}
