package server

import (
	"context"
	"log"
	"time"

	"github.com/qwenyang/xmetau/unigateway/proto/cgi"

	pb "github.com/qwenyang/xmetau/proto/unidao"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func daoUserBarrierLevelToCgi(req *pb.UserBarrierLevel) *cgi.UserBarrierLevel {
	return &cgi.UserBarrierLevel{
		UserId:     req.UserId,
		LevelIndex: req.LevelIndex,
		PassCount:  req.PassCount,
		PassToken:  req.PassToken,
	}
}

func GetUserBarrierLevel(userId uint64) (error, *cgi.UserBarrierLevel) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err, nil
	}
	defer conn.Close()
	c := pb.NewXMetauDaoServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UserBarrierLevelReq{
		UserId: userId,
	}
	r, err := c.GetUserBarrierLevel(ctx, req)
	if err != nil {
		log.Fatalf("asset list failed: %v", err)
		return err, nil
	}

	barrier := daoUserBarrierLevelToCgi(r.Barrier)

	return nil, barrier
}

func UpdateUserBarrierLevel(userId uint64, levelIndex int32, passCount int32, passToken int32) (error, *cgi.UserBarrierLevel) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err, nil
	}
	defer conn.Close()
	c := pb.NewXMetauDaoServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	bar := &pb.UserBarrierLevel{
		UserId:     userId,
		LevelIndex: levelIndex,
		PassCount:  passCount,
		PassToken:  passToken,
	}
	req := &pb.UpdateUserBarrierLevelReq{
		Barrier: bar,
	}
	r, err := c.UpdateUserBarrierLevel(ctx, req)
	if err != nil {
		log.Fatalf("asset list failed: %v", err)
		return err, nil
	}

	return nil, daoUserBarrierLevelToCgi(r.Barrier)
}
