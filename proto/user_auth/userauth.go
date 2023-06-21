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

// Package admin_auth Package main implements a client for Greeter service.
package user_auth

import (
	"audio/apps/app/config"
	pb "audio/proto/pb/cp"
	"context"
	"crypto/tls"
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func AdminAuth(Session string, Action string, Controller string, Param string) bool {
	flag.Parse()
	host := config.GetConfigData()
	ssl := config.GetConfigSsl()

	var conn *grpc.ClientConn
	var err error
	if ssl {
		conn, err = grpc.Dial(host, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		conn, err = grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return false
	}
	defer conn.Close()
	client := pb.NewUserInternalClient(conn)

	resp, err := client.AdminMid(context.Background(), &pb.UserInternalParams_AdminMidReq{
		Session: Session,
	})

	if err != nil {
		log.Fatalf(" %v", err)
		return false
	}

	//mid
	var mid uint32 = resp.GetMid()
	logx.Info(mid)
	if mid > 0 {
		auth, err := client.AdminAuth(context.Background(), &pb.UserInternalParams_AdminAuthReq{
			Action:     Action,
			Controller: Controller,
			Param:      Param,
			Session:    Session,
		})

		if err != nil {
			log.Fatalf(" %v", err)
		}
		return auth.Auth
	}
	return false
}

func ApiAuth(Token string) int {
	flag.Parse()
	host := config.GetConfigData()
	ssl := config.GetConfigSsl()

	var conn *grpc.ClientConn
	var err error
	if ssl {
		conn, err = grpc.Dial(host, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		conn, err = grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	var uid int = 0

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return uid
	}
	defer conn.Close()
	client := pb.NewUserInternalClient(conn)

	resp, err := client.Uid(context.Background(), &pb.UserInternalParams_Token{
		Token: Token,
	})

	if err != nil {
		log.Fatalf(" %v", err)
		return uid
	}

	//mid
	uid = int(resp.Uid)
	return uid
}
