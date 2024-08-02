package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/qwenyang/xmetau/unigateway/proto/cgi"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime" // 注意v2版本s
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", "0.0.0.0:9990")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册Greeter service到server
	cgi.RegisterXMetauCgiSvrServer(s, NewCgiProxyServiceService())
	// 8080端口启动gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:9990")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:9990",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// 注册Greeter
	err = cgi.RegisterXMetauCgiSvrHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    "0.0.0.0:9991",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:9991")
	log.Fatalln(gwServer.ListenAndServe())
}
