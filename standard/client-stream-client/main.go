/*
 * based on
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	pb "github.com/bakins/grpc-the-hard-way/services/greetings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

const (
	address = "localhost:50051"
)

func main() {
	useGzip := flag.Bool("gzip", false, "enable gzip compression")
	flag.Parse()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	if *useGzip {
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.CrowdGreeting(ctx)
	if err != nil {
		log.Fatalf("could send request: %v", err)
	}

	go sendNames(stream)

	// will block until response has been sent
	// or timeout is done
	var resp pb.GreetingReply
	err = stream.RecvMsg(&resp)
	if err != nil {
		log.Fatalf("could not receive response: %v", err)
	}

	log.Printf("Greeting: %s", resp.Message)
}

func sendNames(stream pb.Greeter_CrowdGreetingClient) {
	// send 10 "names"
	for i := 0; i < 10; i++ {
		req := pb.GreetingRequest{
			Name: strconv.Itoa(i),
		}

		if err := stream.Send(&req); err != nil {
			// in a "real" program, we would handle this error, but for now
			// just exit
			log.Fatalf("failed to send request: %v", err)
		}
	}

	// close to signify that we are done sending requests
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("failed to close stream: %v", err)
	}
}
