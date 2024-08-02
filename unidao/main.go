package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	common "github.com/qwenyang/xmetau/proto/common"
	pb "github.com/qwenyang/xmetau/proto/unidao"
	tables "github.com/qwenyang/xmetau/unidao/tables"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 60003, "The server port")
)

// server is used to implement XMetauDaoService.
type XMetauDaoSvr struct {
	pb.UnimplementedXMetauDaoServiceServer
}

// Login implements XMetauDaoService
func (s *XMetauDaoSvr) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	log.Printf("Login Received: %v", in.Login)
	err, isNewUser, account := tables.LoginRegister(in.Login)
	if err != nil {
		return nil, err
	}
	// 默认创建一个用户基础属性信息
	if isNewUser {
		ua := &pb.UserAttribute{
			UserId:         account.UserId,
			NickName:       "",
			AvatarUrl:      "",
			NoviceTraining: 0,
			PlayLevel:      0,
			GoldCoin:       10000,
			WinNum:         0,
			LoseNum:        0,
			TieNum:         0,
			GameName:       in.GameName,
			ModifyTime:     "",
		}
		tables.UpdateUserTable(ua)
	}
	resp := &pb.LoginResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Login:   in.Login,
		Account: account,
	}
	return resp, nil
}

// QueryUser implements XMetauDaoService
func (s *XMetauDaoSvr) QueryUser(ctx context.Context, in *pb.QueryUserReq) (*pb.QueryUserResp, error) {
	log.Printf("Query User Received: %v", in)
	err, userResp := tables.GetUserTable(in.UserId)
	if err != nil {
		return nil, err
	}
	resp := &pb.QueryUserResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		User: userResp,
	}
	return resp, nil
}

// UpdateUserHeader implements XMetauDaoService
func (s *XMetauDaoSvr) UpdateUserHeader(ctx context.Context, in *pb.UpdateUserHeaderReq) (*pb.UpdateUserResp, error) {
	log.Printf("Update User Header Received: %v", in)
	err, userResp := tables.UpdateUserHeader(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UpdateUserResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		User: userResp,
	}
	return resp, nil
}

// UpdateUserNoviceTraining implements XMetauDaoService
func (s *XMetauDaoSvr) UpdateUserNoviceTraining(ctx context.Context, in *pb.UpdateUserTrainingReq) (*pb.UpdateUserResp, error) {
	log.Printf("Update User Training Received: %v", in)
	err, userResp := tables.UpdateUserNoviceTraining(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UpdateUserResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		User: userResp,
	}
	return resp, nil
}

// UpdateUserHeader implements XMetauDaoService
func (s *XMetauDaoSvr) UpdateUserPlayLevel(ctx context.Context, in *pb.UpdateUserLevelReq) (*pb.UpdateUserResp, error) {
	log.Printf("Update User Level Received: %v", in)
	err, userResp := tables.UpdateUserPlayLevel(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UpdateUserResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		User: userResp,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) UpdateUserCoin(ctx context.Context, in *pb.UpdateUserCoinReq) (*pb.UpdateUserResp, error) {
	log.Printf("Update User coin Received: %v", in)
	err, userResp := tables.UpdateUserCoin(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UpdateUserResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		User: userResp,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) UpdateGameNum(ctx context.Context, in *pb.UpdateGameNumReq) (*pb.UpdateUserResp, error) {
	log.Printf("Update User game num Received: %v", in)
	err, userResp := tables.UpdateGameNum(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UpdateUserResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		User: userResp,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) QueryUserRankList(ctx context.Context, in *pb.UserListReq) (*pb.UserListResp, error) {
	log.Printf("Update User rank list Received: %v", in)
	err, pageData, userList := tables.QueryUserRankList(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UserListResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		PageData: pageData,
		UserList: userList,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) QueryRobotUserList(ctx context.Context, in *pb.UserListReq) (*pb.UserListResp, error) {
	log.Printf("Update Robot user list Received: %v", in)
	err, pageData, userList := tables.QueryRobotUserList(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UserListResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		PageData: pageData,
		UserList: userList,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) AssetList(ctx context.Context, in *pb.CommonUserReq) (*pb.AssetListResp, error) {
	log.Printf("AssetList Received: %v", in)
	err, assetList := tables.AssetList(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.AssetListResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Assets: assetList,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) UserAssetList(ctx context.Context, in *pb.CommonUserReq) (*pb.UserAssetResp, error) {
	log.Printf("UserAssetList Received: %v", in)
	err, userAssetList := tables.UserAssetList(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UserAssetResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		UserAssets: userAssetList,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) UpdateUserAsset(ctx context.Context, in *pb.UpdateUserAssetReq) (*pb.UserAssetResp, error) {
	log.Printf("UpdateUserAsset Received: %v", in)
	err, userAssetList := tables.UpdateUserAsset(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UserAssetResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		UserAssets: userAssetList,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) SettingList(ctx context.Context, in *pb.GameSettingReq) (*pb.GameSettingResp, error) {
	log.Printf("SettingList Received: %v", in)
	err, settings := tables.SettingList(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.GameSettingResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Settings: settings,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) UpdateSetting(ctx context.Context, in *pb.UpdateSettingReq) (*pb.UpdateSettingResp, error) {
	log.Printf("UpdateSetting Received: %v", in)
	err, settings := tables.UpdateSetting(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UpdateSettingResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Settings: settings,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) GetTeamInfo(ctx context.Context, in *pb.TeamInfoReq) (*pb.TeamInfoResp, error) {
	log.Printf("GetTeamInfo Received: %v", in)
	err, team := tables.GetTeamInfo(in.TeamId)
	if err != nil {
		return nil, err
	}
	resp := &pb.TeamInfoResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Team: team,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) GetUserTeam(ctx context.Context, in *pb.UserTeamReq) (*pb.UserTeamResp, error) {
	log.Printf("GetUserTeam Received: %v", in)
	err, team := tables.GetUserTeam(in.UserId)
	if err != nil {
		return nil, err
	}
	resp := &pb.UserTeamResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Team: team,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) CreateTeam(ctx context.Context, in *pb.CreateTeamReq) (*pb.CreateTeamResp, error) {
	log.Printf("CreateTeam Received: %v", in)
	err, team := tables.CreateTeam(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.CreateTeamResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Team: team,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) JoinTeam(ctx context.Context, in *pb.JoinTeamReq) (*pb.JoinTeamResp, error) {
	log.Printf("JoinTeam Received: %v", in)
	err, team := tables.JoinTeam(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.JoinTeamResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Team: team,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) GetTeamRankList(ctx context.Context, in *pb.TeamRankListReq) (*pb.TeamRankListResp, error) {
	log.Printf("GetTeamRankList Received: %v", in)
	err, teams := tables.GetTeamRankList(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.TeamRankListResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Teams: teams,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) UpdateUserTeam(ctx context.Context, in *pb.UpdateUserTeamReq) (*pb.UpdateUserTeamResp, error) {
	log.Printf("UpdateUserTeam Received: %v", in)
	err, team := tables.UpdateUserTeam(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UpdateUserTeamResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Team: team,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) GetUserBarrierLevel(ctx context.Context, in *pb.UserBarrierLevelReq) (*pb.UserBarrierLevelResp, error) {
	log.Printf("GetUserBarrierLevel Received: %v", in)
	err, barrier := tables.GetUserBarrierLevel(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UserBarrierLevelResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Barrier: barrier,
	}
	return resp, nil
}

func (s *XMetauDaoSvr) UpdateUserBarrierLevel(ctx context.Context, in *pb.UpdateUserBarrierLevelReq) (*pb.UpdateUserBarrierLevelResp, error) {
	log.Printf("UpdateUserBarrierLevel Received: %v", in)
	err, barrier := tables.UpdateUserBarrierLevel(in)
	if err != nil {
		return nil, err
	}
	resp := &pb.UpdateUserBarrierLevelResp{
		Header: &common.Header{
			Code:    0,
			Message: "ok",
		},
		Barrier: barrier,
	}
	return resp, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterXMetauDaoServiceServer(s, &XMetauDaoSvr{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
