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
	"io"
	"log"
	"time"

	pb "github.com/bakins/grpc-the-hard-way/services/greetings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
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

	// Contact the server and print out its response.
	name := defaultName
	args := flag.Args()
	if len(args) > 1 {
		name = args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stream, err := c.ShareGreetings(ctx, &pb.GreetingRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

READ:
	for {
		r, err := stream.Recv()
		switch err {
		case nil:
		case io.EOF:
			break READ
		default:
			log.Fatalf("failed to recv: %v", err)
		}
		log.Printf("Greeting: %s", r.Message)
	}
}
