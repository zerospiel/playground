/*
 *
 * Copyright 2018 gRPC authors.
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
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	// channelz "github.com/rantav/go-grpc-channelz"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	// channelzservice "google.golang.org/grpc/channelz/service"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

const (
	defaultName = "world"
)

func main() {
	/***** Set up the server serving channelz service. *****/
	const grpcBindAddress = "127.0.0.1:50050"
	lis, err := net.Listen("tcp4", grpcBindAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// channelzservice.RegisterChannelzServiceToServer(s)
	reflection.Register(s)
	go s.Serve(lis)
	defer s.Stop()

	adminLis, err := net.Listen("tcp4", "127.0.0.1:8082")
	if err != nil {
		panic(fmt.Errorf("failed to listen admin: %w", err))
	}
	defer adminLis.Close()

	// http.Handle("/", channelz.CreateHandler("/", grpcBindAddress))
	go http.Serve(adminLis, nil)

	/***** Initialize manual resolver and Dial *****/
	r := manual.NewBuilderWithScheme("ozon")
	resolver.Register(r)
	// defer resolver.UnregisterForTesting("ozon")
	// Set up a connection to the server.
	conn, err := grpc.Dial(
		r.Scheme()+":///test.server",
		grpc.WithAuthority(""),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [ { "round_robin": {} } ]}`),
	)
	if err != nil {
		panic(fmt.Errorf("did not connect: %w", err))
	}
	defer conn.Close()
	// Manually provide resolved addresses for the target.
	r.UpdateState(resolver.State{
		Addresses: []resolver.Address{{Addr: ":10001"}, {Addr: ":10002"}, {Addr: ":10003"}},
	})

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// /***** Make 100 SayHello RPCs *****/
	for i := 0; i < 100; i++ {
		// Setting a 150ms timeout on the RPC.
		ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
		defer cancel()
		p := peer.Peer{}
		resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: name}, grpc.Peer(&p))
		if err != nil {
			// conn.ResetConnectBackoff()
			// GRPC_GO_LOG_SEVERITY_LEVEL=INFO GRPC_GO_LOG_VERBOSITY_LEVEL=99
			log.Printf("could not greet %s: %v", p.Addr.String(), err)
		} else {
			log.Printf("Greeting %s: %s", p.Addr.String(), resp.Message)
		}
	}

	/***** Wait for CTRL+C to exit *****/
	// Unless you exit the program with CTRL+C, channelz data will be available for querying.
	// Users can take time to examine and learn about the info provided by channelz.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	// Block until a signal is received.
	<-ch
}
