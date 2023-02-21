/*
 *
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
	"log"
	"testing"
)

var (
	addr             = "localhost:7000"
	tenantID         = "demo123821"
	tenantConfigFile = "conf/tiflash-templete.toml"
)

//func InitGrpcServer() {
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//	pb.RegisterAssignServer(s, &server{})
//	log.Printf("server listening at %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//	go TiFlashMaintainer()
//}
//
//func TestAssignAndUnassignTenant(t *testing.T) {
//	go InitGrpcServer()
//	flag.Parse()
//	// Set up a connection to the server.
//	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := pb.NewAssignClient(conn)
//
//	// Contact the server and print out its response.
//	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
//	defer cancel()
//	r, err := c.AssignTenant(ctx, &pb.AssignRequest{TenantID: tenantID, TidbStatusAddr: "123.123.123.123:1000", PdAddr: "179.1.1.1:2000"})
//	if err != nil {
//		log.Fatalf("could not assign: %v", err)
//	}
//	if r.HasErr {
//		log.Fatalf("assign failed: %v", r.ErrInfo)
//	} else {
//		log.Printf("assign succeeded")
//	}
//
//}

//func TestGetCurrentTenant(t *testing.T) {
//	flag.Parse()
//	// Set up a connection to the server.
//	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := pb.NewAssignClient(conn)
//
//	// Contact the server and print out its response.
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	r, err := c.GetCurrentTenant(ctx, &emptypb.Empty{})
//	log.Printf("current tenant: %v", r.TenantID)
//}
//
//func TestUnassignTenant(t *testing.T) {
//	flag.Parse()
//	// Set up a connection to the server.
//	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := pb.NewAssignClient(conn)
//
//	// Contact the server and print out its response.
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	r, err := c.UnassignTenant(ctx, &pb.UnassignRequest{AssertTenantID: tenantID})
//	if err != nil {
//		log.Fatalf("could not unassign: %v", err)
//	}
//	if r.HasErr {
//		log.Fatalf("unassign failed: %v", r.ErrInfo)
//	} else {
//		log.Printf("unassign succeeded")
//	}
//}

func TestInitTiFlashConf(t *testing.T) {
	err := InitTiFlashConf()
	if err != nil {
		log.Fatalf("init tiflash conf failed: %v", err)
	}
	err = RenderTiFlashConf("conf/tiflash.toml", "123.123.123.123:1000", "179.1.1.1:2000", "")
	if err != nil {
		log.Fatalf("RenderTiFlashConf failed: %v", err)
	}
}

func TestInitTiFlashConf2(t *testing.T) {
	err := InitTiFlashConf()
	if err != nil {
		log.Fatalf("init tiflash conf failed: %v", err)
	}
	err = RenderTiFlashConf("conf/tiflash.toml", "123.123.123.123:1000", "179.1.1.1:2000", "fixpool-use-autoscaler-false")
	if err != nil {
		log.Fatalf("RenderTiFlashConf failed: %v", err)
	}
}

func TestInitTiFlashConf3(t *testing.T) {
	err := InitTiFlashConf()
	if err != nil {
		log.Fatalf("init tiflash conf failed: %v", err)
	}
	err = RenderTiFlashConf("conf/tiflash.toml", "123.123.123.123:1000", "179.1.1.1:2000", "fixpool-use-autoscaler-true")
	if err != nil {
		log.Fatalf("RenderTiFlashConf failed: %v", err)
	}
}
