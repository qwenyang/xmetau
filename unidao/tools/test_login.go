package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/qwenyang/xmetau/proto/unidao"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addrLogin = flag.String("addr", "localhost:60003", "the address to connect to")
	appId     = flag.String("appId", "123456", "appid to set")
	openId    = flag.String("openId", "uuuuuabcd1234xyz", "open id to set")
	unionId   = flag.String("unionId", "abcd1234xyz", "union id to set")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addrLogin, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewXMetauDaoServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	login := &pb.Login{
		LoginType: 0,
		AppId:     *appId,
		OpenId:    *openId,
		UnionId:   *unionId,
	}

	req := &pb.LoginReq{
		Login: login,
	}
	r, err := c.Login(ctx, req)
	if err != nil {
		log.Fatalf("user login failed: %v", err)
	}
	log.Printf("login ok %v", r)
}
