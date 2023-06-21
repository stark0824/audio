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

// Package client Package main implements a client for Greeter service.
package user_auth_test

import (
	pb "audio/proto/pb/cp"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial("testing.gongzicp.com:1443", grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserInternalClient(conn)

	resp, err := client.AdminMid(context.Background(), &pb.UserInternalParams_AdminMidReq{
		Session: "b1efbc6352r4jgif5qku3l35km",
	})

	if err != nil {
		log.Fatalf(" %v", err)
	}

	//mid
	var mid uint32 = resp.GetMid()
	if mid > 0 {
		auth, err := client.AdminAuth(context.Background(), &pb.UserInternalParams_AdminAuthReq{
			Action:     "dramaList",
			Controller: "audio",
			Param:      "",
			Session:    "b1efbc6352r4jgif5qku3l35km",
		})

		if err != nil {
			log.Fatalf(" %v", err)
		}

		fmt.Printf("%t", auth.Auth)
	}
}
