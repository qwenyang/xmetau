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

func daoGameSettingToCgi(req *pb.GameSetting) *cgi.GameSetting {
	return &cgi.GameSetting{
		SetId:    req.SetId,
		SetType:  req.SetType,
		SetKey:   req.SetKey,
		SetValue: req.SetValue,
	}
}

func SettingList(userId uint64, setId uint64, setType int32) (error, []*cgi.GameSetting) {
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

	req := &pb.GameSettingReq{
		UserId:  userId,
		SetId:   setId,
		SetType: setType,
	}
	r, err := c.SettingList(ctx, req)
	if err != nil {
		log.Fatalf("asset list failed: %v", err)
		return err, nil
	}

	settingList := make([]*cgi.GameSetting, len(r.Settings))
	for i := 0; i < len(r.Settings); i++ {
		settingList[i] = daoGameSettingToCgi(r.Settings[i])
	}

	return nil, settingList
}

func UpdateSetting(userId uint64, setId uint64, setType int32, setKey string, setValue string) (error, []*cgi.GameSetting) {
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

	req := &pb.UpdateSettingReq{
		UserId:   userId,
		SetId:    setId,
		SetType:  setType,
		SetKey:   setKey,
		SetValue: setValue,
	}
	r, err := c.UpdateSetting(ctx, req)
	if err != nil {
		log.Fatalf("asset list failed: %v", err)
		return err, nil
	}

	settingList := make([]*cgi.GameSetting, len(r.Settings))
	for i := 0; i < len(r.Settings); i++ {
		settingList[i] = daoGameSettingToCgi(r.Settings[i])
	}

	return nil, settingList
}
