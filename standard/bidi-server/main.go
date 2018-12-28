/*
 * based on examples by
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
	"io"
	"log"
	"math/rand"
	"net"

	pb "github.com/bakins/grpc-the-hard-way/services/greetings"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"
)

const (
	port = "127.0.0.1:50051"
)

type server struct{}

func (s *server) ShareGreetings(req *pb.GreetingRequest, stream pb.Greeter_ShareGreetingsServer) error {
	// unused in this example, but defined to implement the interface
	return nil
}

func (s *server) CrowdGreeting(stream pb.Greeter_CrowdGreetingServer) error {
	// unused in this example, but needed to implement the interface
	return nil
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

func (s *server) StreamGreetings(stream pb.Greeter_StreamGreetingsServer) error {
READ:
	for {
		req, err := stream.Recv()
		switch err {
		case nil:
			// keep reading
		case io.EOF:
			// done
			break READ
		default:
			log.Printf("failed to read from stream: %v", err)
			return err
		}

		n := rand.Int() % len(greetings)

		resp := pb.GreetingReply{
			Message: greetings[n] + " " + req.GetName(),
		}

		if err := stream.Send(&resp); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
