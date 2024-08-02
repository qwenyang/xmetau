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

func daoUserToCgiUser(userInfo *pb.UserAttribute) *cgi.UserAttribute {
	return &cgi.UserAttribute{
		UserId:         userInfo.UserId,
		NickName:       userInfo.NickName,
		AvatarUrl:      userInfo.AvatarUrl,
		NoviceTraining: userInfo.NoviceTraining,
		PlayLevel:      userInfo.PlayLevel,
		GoldCoin:       userInfo.GoldCoin,
		WinNum:         userInfo.WinNum,
		LoseNum:        userInfo.LoseNum,
		TieNum:         userInfo.TieNum,
		ModifyTime:     userInfo.ModifyTime,
		CreateTime:     userInfo.CreateTime,
	}
}

func GetUserAttribute(userId uint64) (error, *cgi.UserAttribute) {
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

	req := &pb.QueryUserReq{
		UserId: userId,
	}
	r, err := c.QueryUser(ctx, req)
	if err != nil {
		log.Fatalf("query user failed: %v", err)
		return err, nil
	}

	user := daoUserToCgiUser(r.User)
	return nil, user
}

func UpdateUserHeader(in *cgi.UpdateUserHeaderReq) (error, *cgi.UserAttribute) {
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

	req := &pb.UpdateUserHeaderReq{
		UserId:    in.UserId,
		NickName:  in.NickName,
		AvatarUrl: in.AvatarUrl,
	}
	r, err := c.UpdateUserHeader(ctx, req)
	if err != nil {
		log.Fatalf("update user failed: %v", err)
		return err, nil
	}
	return nil, daoUserToCgiUser(r.User)
}

func UpdateUserNoviceTraining(in *cgi.UpdateUserTrainingReq) (error, *cgi.UserAttribute) {
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

	req := &pb.UpdateUserTrainingReq{
		UserId:         in.UserId,
		NoviceTraining: in.NoviceTraining,
	}
	r, err := c.UpdateUserNoviceTraining(ctx, req)
	if err != nil {
		log.Fatalf("update user failed: %v", err)
		return err, nil
	}
	return nil, daoUserToCgiUser(r.User)
}

func UpdateUserPlayLevel(in *cgi.UpdateUserLevelReq) (error, *cgi.UserAttribute) {
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

	req := &pb.UpdateUserLevelReq{
		UserId:    in.UserId,
		PlayLevel: in.PlayLevel,
		DiffLevel: in.DiffLevel,
	}
	r, err := c.UpdateUserPlayLevel(ctx, req)
	if err != nil {
		log.Fatalf("update user failed: %v", err)
		return err, nil
	}
	return nil, daoUserToCgiUser(r.User)
}

func UpdateUserCoin(in *cgi.UpdateUserCoinReq) (error, *cgi.UserAttribute) {
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

	req := &pb.UpdateUserCoinReq{
		UserId:   in.UserId,
		GoldCoin: in.GoldCoin,
		DiffCoin: in.DiffCoin,
	}
	r, err := c.UpdateUserCoin(ctx, req)
	if err != nil {
		log.Fatalf("update user failed: %v", err)
		return err, nil
	}
	return nil, daoUserToCgiUser(r.User)
}

func UpdateGameNum(in *cgi.UpdateGameNumReq) (error, *cgi.UserAttribute) {
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

	req := &pb.UpdateGameNumReq{
		UserId:  in.UserId,
		WinNum:  in.WinNum,
		LoseNum: in.LoseNum,
		TieNum:  in.TieNum,
		DiffNum: in.DiffNum,
	}
	r, err := c.UpdateGameNum(ctx, req)
	if err != nil {
		log.Fatalf("update user failed: %v", err)
		return err, nil
	}
	return nil, daoUserToCgiUser(r.User)
}

func QueryUserRankList(in *cgi.UserListReq) (error, *cgi.PageData, []*cgi.UserAttribute) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err, nil, nil
	}
	defer conn.Close()
	c := pb.NewXMetauDaoServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UserListReq{
		UserId:    in.UserId,
		PageSize:  in.PageSize,
		PageIndex: in.PageIndex,
		GameName:  in.GameName,
	}
	r, err := c.QueryUserRankList(ctx, req)
	if err != nil {
		log.Fatalf("update user failed: %v", err)
		return err, nil, nil
	}
	pageData := &cgi.PageData{
		Total:     r.PageData.Total,
		PageSize:  r.PageData.PageSize,
		PageIndex: r.PageData.PageIndex,
	}
	userList := make([]*cgi.UserAttribute, len(r.UserList))
	for i := 0; i < len(r.UserList); i++ {
		userList[i] = daoUserToCgiUser(r.UserList[i])
	}
	return nil, pageData, userList
}

func QueryRobotUserList(in *cgi.UserListReq) (error, *cgi.PageData, []*cgi.UserAttribute) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err, nil, nil
	}
	defer conn.Close()
	c := pb.NewXMetauDaoServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UserListReq{
		UserId:    in.UserId,
		PageSize:  in.PageSize,
		PageIndex: in.PageIndex,
	}
	r, err := c.QueryRobotUserList(ctx, req)
	if err != nil {
		log.Fatalf("update user failed: %v", err)
		return err, nil, nil
	}
	pageData := &cgi.PageData{
		Total:     r.PageData.Total,
		PageSize:  r.PageData.PageSize,
		PageIndex: r.PageData.PageIndex,
	}
	userList := make([]*cgi.UserAttribute, len(r.UserList))
	for i := 0; i < len(r.UserList); i++ {
		userList[i] = daoUserToCgiUser(r.UserList[i])
	}
	return nil, pageData, userList
}
