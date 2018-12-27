package main

import (
	"encoding/binary"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	pb "github.com/bakins/grpc-the-hard-way/services/helloworld"
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
		log.Fatalf("failed to start server: %v", err)
	}
}

func handleSayHello(w http.ResponseWriter, r *http.Request) {
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md
	// The repeated sequence of Length-Prefixed-Message items is delivered in DATA frames
	// Length-Prefixed-Message → Compressed-Flag Message-Length Message
	// Compressed-Flag → 0 / 1 # encoded as 1 byte unsigned integer
	// Message-Length → {length of Message} # encoded as 4 byte unsigned integer (big endian)
	// Message → *{binary octet}

	// first byte is compressed flag, next 4 are an unsigned integer
	prefix := []byte{0, 0, 0, 0, 0}

	_, err := r.Body.Read(prefix)
	if err != nil {
		http.Error(w, "failed to read prefix: "+err.Error(), http.StatusBadRequest)
		return
	}

	// this version does not support compression, so ignore the
	// flag

	// determine the length of the message.  Future versions
	// should ensure this is a valid length - ie, not 0 and not greater
	// than a configured maximum size
	length := binary.BigEndian.Uint32(prefix[1:])

	// now read the message
	message := make([]byte, length)

	_, err = r.Body.Read(message)
	if err != nil {
		http.Error(w, "failed to read message: "+err.Error(), http.StatusBadRequest)
		return
	}

	//message contains a marshalled protobuf
	var req pb.HelloRequest
	if err := proto.Unmarshal(message, &req); err != nil {
		http.Error(w, "failed to unmarshal message: "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := pb.HelloReply{
		Message: "hello " + req.GetName(),
	}

	// marshal our response
	body, err := proto.Marshal(&resp)
	if err != nil {
		http.Error(w, "failed to marshal response body: "+err.Error(), http.StatusBadRequest)
		return
	}
	// create a length prefix for our response.
	// We do not support compression in this version
	prefix[0] = 0
	binary.BigEndian.PutUint32(prefix[1:], uint32(len(body)))

	// Set correct content type
	w.Header().Set("Content-Type", "application/grpc+proto")

	// We must include trailers for status and message.
	// See https://golang.org/pkg/net/http/#example_ResponseWriter_trailers
	w.Header().Set("Trailer", "grpc-status, grpc-message")

	// any non-protocol error should still return an HTTP 200
	// and use grpc-status to report the error
	w.WriteHeader(200)

	// write the prefix
	_, err = w.Write(prefix)
	if err != nil {
		// we have already written the headers, so just log error
		log.Printf("failed to write prefix: %v", err)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		log.Printf("failed to write body: %v", err)
		return
	}

	w.Header().Set("grpc-status", strconv.Itoa(0))
	w.Header().Set("grpc-message", "ok")
}
