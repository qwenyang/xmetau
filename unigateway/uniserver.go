package main

import (
	"context"
	"log"

	"github.com/qwenyang/xmetau/unigateway/proto/cgi"

	proxy "github.com/qwenyang/xmetau/unigateway/server"
)

type cgiProxyService struct {
}

func NewCgiProxyServiceService() *cgiProxyService {
	return &cgiProxyService{}
}

// 更新用户头像信息
func (h *cgiProxyService) UpdateUserHeader(ctx context.Context, in *cgi.UpdateUserHeaderReq) (*cgi.UpdateUserResp, error) {
	log.Printf("update user header %v", in)
	err, user := proxy.UpdateUserHeader(in)
	if err != nil {
		return nil, err
	}

	resp := &cgi.UpdateUserResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		User: user,
	}
	return resp, err
}

// 更新用户新手训练信息
func (h *cgiProxyService) UpdateUserNoviceTraining(ctx context.Context, in *cgi.UpdateUserTrainingReq) (*cgi.UpdateUserResp, error) {
	err, user := proxy.UpdateUserNoviceTraining(in)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UpdateUserResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		User: user,
	}
	return resp, err
}

// 更新用户等级信息
func (h *cgiProxyService) UpdateUserPlayLevel(ctx context.Context, in *cgi.UpdateUserLevelReq) (*cgi.UpdateUserResp, error) {
	log.Printf("update user play level: %v", in)
	err, user := proxy.UpdateUserPlayLevel(in)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UpdateUserResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		User: user,
	}
	return resp, err
}

func (h *cgiProxyService) UpdateUserCoin(ctx context.Context, in *cgi.UpdateUserCoinReq) (*cgi.UpdateUserResp, error) {
	log.Printf("update user coin: %v", in)
	err, user := proxy.UpdateUserCoin(in)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UpdateUserResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		User: user,
	}
	return resp, err
}

func (h *cgiProxyService) UpdateGameNum(ctx context.Context, in *cgi.UpdateGameNumReq) (*cgi.UpdateUserResp, error) {
	log.Printf("update user num: %v", in)
	err, user := proxy.UpdateGameNum(in)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UpdateUserResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		User: user,
	}
	return resp, err
}

func (h *cgiProxyService) QueryUserRankList(ctx context.Context, in *cgi.UserListReq) (*cgi.UserListResp, error) {
	log.Printf("rank list req: %s %v %+v %#v ", in.GetGameName(), in, in, in)
	err, pageData, userList := proxy.QueryUserRankList(in)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UserListResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		PageData: pageData,
		UserList: userList,
	}
	return resp, err
}

func (h *cgiProxyService) QueryRobotUserList(ctx context.Context, in *cgi.UserListReq) (*cgi.UserListResp, error) {
	err, pageData, userList := proxy.QueryRobotUserList(in)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UserListResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		PageData: pageData,
		UserList: userList,
	}
	return resp, err
}

// 查询用户信息
func (h *cgiProxyService) QueryUser(ctx context.Context, in *cgi.QueryUserReq) (*cgi.QueryUserResp, error) {
	log.Printf("query user attribute: %v", in)
	err, user := proxy.GetUserAttribute(in.UserId)
	if err != nil {
		return nil, err
	}
	log.Printf("query user attribute: %v", user)

	return &cgi.QueryUserResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		User: user,
	}, nil
}

func (h *cgiProxyService) AssetList(ctx context.Context, in *cgi.CommonUserReq) (*cgi.AssetListResp, error) {
	err, assets := proxy.AssetList(in.UserId)
	if err != nil {
		return nil, err
	}
	resp := &cgi.AssetListResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		Assets: assets,
	}
	return resp, nil
}

func (h *cgiProxyService) UserAssetList(ctx context.Context, in *cgi.CommonUserReq) (*cgi.UserAssetResp, error) {
	err, userAssets := proxy.UserAssetList(in.UserId)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UserAssetResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		UserAssets: userAssets,
	}
	return resp, nil
}

func (h *cgiProxyService) UpdateUserAsset(ctx context.Context, in *cgi.UpdateUserAssetReq) (*cgi.UserAssetResp, error) {
	err, userAssets := proxy.UpdateUserAsset(in.UserId, in.AssetId, in.ExpirationTime)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UserAssetResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		UserAssets: userAssets,
	}
	return resp, nil
}

func (h *cgiProxyService) SettingList(ctx context.Context, in *cgi.GameSettingReq) (*cgi.GameSettingResp, error) {
	err, settings := proxy.SettingList(in.UserId, in.SetId, in.SetType)
	if err != nil {
		return nil, err
	}
	resp := &cgi.GameSettingResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		Settings: settings,
	}
	return resp, nil
}

func (h *cgiProxyService) UpdateSetting(ctx context.Context, in *cgi.UpdateSettingReq) (*cgi.UpdateSettingResp, error) {
	log.Printf("UpdateSetting: %v", in)
	err, settings := proxy.UpdateSetting(in.UserId, in.SetId, in.SetType, in.SetKey, in.SetValue)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UpdateSettingResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		Settings: settings,
	}
	return resp, nil
}

func (h *cgiProxyService) GetTeamInfo(ctx context.Context, in *cgi.TeamInfoReq) (*cgi.TeamInfoResp, error) {
	log.Printf("GetUserTeam: %v", in)
	err, team := proxy.GetTeamInfo(in.UserId, in.TeamId)
	if err != nil {
		return nil, err
	}
	resp := &cgi.TeamInfoResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		Team: team,
	}
	return resp, nil
}

func (h *cgiProxyService) GetUserTeam(ctx context.Context, in *cgi.UserTeamReq) (*cgi.UserTeamResp, error) {
	log.Printf("GetUserTeam: %v", in)
	err, team := proxy.GetUserTeam(in.UserId)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UserTeamResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		Team: team,
	}
	return resp, nil
}

func (h *cgiProxyService) CreateTeam(ctx context.Context, in *cgi.CreateTeamReq) (*cgi.CreateTeamResp, error) {
	log.Printf("CreateTeam: %v", in)
	err, team := proxy.CreateTeam(in.UserId, in.TeamId, in.TeamName, in.CharType)
	if err != nil {
		return nil, err
	}
	resp := &cgi.CreateTeamResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		Team: team,
	}
	return resp, nil
}

func (h *cgiProxyService) JoinTeam(ctx context.Context, in *cgi.JoinTeamReq) (*cgi.JoinTeamResp, error) {
	log.Printf("JoinTeam: %v", in)
	err, team := proxy.JoinTeam(in.UserId, in.TeamId, in.CharType)
	if err != nil {
		return nil, err
	}
	resp := &cgi.JoinTeamResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		Team: team,
	}
	return resp, nil
}

func (h *cgiProxyService) GetTeamRankList(ctx context.Context, in *cgi.TeamRankListReq) (*cgi.TeamRankListResp, error) {
	log.Printf("GetTeamRankList: %v", in)
	err, teams := proxy.GetTeamRankList(in.UserId)
	if err != nil {
		return nil, err
	}
	resp := &cgi.TeamRankListResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		Teams: teams,
	}
	return resp, nil
}

func (h *cgiProxyService) UpdateUserTeam(ctx context.Context, in *cgi.UpdateUserTeamReq) (*cgi.UpdateUserTeamResp, error) {
	log.Printf("UpdateUserTeam: %v", in)
	err, team := proxy.UpdateUserTeam(in)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UpdateUserTeamResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		Team: team,
	}
	return resp, nil
}

func (h *cgiProxyService) GetUserBarrierLevel(ctx context.Context, in *cgi.UserBarrierLevelReq) (*cgi.UserBarrierLevelResp, error) {
	log.Printf("GetUserBarrierLevel: %v", in)
	err, barrier := proxy.GetUserBarrierLevel(in.UserId)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UserBarrierLevelResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		Barrier: barrier,
	}
	return resp, nil
}

func (h *cgiProxyService) UpdateUserBarrierLevel(ctx context.Context, in *cgi.UpdateUserBarrierLevelReq) (*cgi.UpdateUserBarrierLevelResp, error) {
	log.Printf("UpdateUserBarrierLevel: %v", in)
	err, barrier := proxy.UpdateUserBarrierLevel(in.UserId, in.LevelIndex, in.PassCount, in.PassToken)
	if err != nil {
		return nil, err
	}
	resp := &cgi.UpdateUserBarrierLevelResp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		Barrier: barrier,
	}
	return resp, nil
}
