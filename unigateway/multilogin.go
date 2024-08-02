package main

import (
	"context"
	"log"

	"github.com/qwenyang/xmetau/unigateway/proto/cgi"
	proxy "github.com/qwenyang/xmetau/unigateway/server"
	platform "github.com/qwenyang/xmetau/unigateway/server/platform"
)

// 登录接口，也会返回用户信息 => 默认是微信小游戏斗子象棋的Login
func (h *cgiProxyService) WxChessLogin(ctx context.Context, in *cgi.LoginReq) (*cgi.LoginRsp, error) {
	log.Printf("login %v", in)
	err, session := platform.WxCode2Session(in.Code, WxChessAppId, WxChessAppSecret)
	if err != nil {
		return nil, err
	}
	gameName := "douzi"
	err, userId := proxy.LoginGetUserId(session, WxChessAppId, gameName)
	if err != nil {
		return nil, err
	}
	err, user := proxy.GetUserAttribute(userId)
	if err != nil {
		return nil, err
	}
	resp := &cgi.LoginRsp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		CodeSession: session,
		User:        user,
	}
	return resp, err
}

// 微信桌球登录接口，也会返回用户信息
func (h *cgiProxyService) WxBilliardLogin(ctx context.Context, in *cgi.LoginReq) (*cgi.LoginRsp, error) {
	log.Printf("billiard wx login %v", in)
	err, session := platform.WxCode2Session(in.Code, WxBilliardAppId, WxBilliardAppSecret)
	if err != nil {
		return nil, err
	}
	gameName := "billiard"
	err, userId := proxy.LoginGetUserId(session, WxBilliardAppId, gameName)
	if err != nil {
		return nil, err
	}
	err, user := proxy.GetUserAttribute(userId)
	if err != nil {
		return nil, err
	}
	resp := &cgi.LoginRsp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		CodeSession: session,
		User:        user,
	}
	return resp, err
}

// 微信坦克登录接口，也会返回用户信息
func (h *cgiProxyService) WxTankLogin(ctx context.Context, in *cgi.LoginReq) (*cgi.LoginRsp, error) {
	log.Printf("billiard wx login %v", in)
	err, session := platform.WxCode2Session(in.Code, WxTankAppId, WxTankAppSecret)
	if err != nil {
		return nil, err
	}
	gameName := "tank"
	err, userId := proxy.LoginGetUserId(session, WxTankAppId, gameName)
	if err != nil {
		return nil, err
	}
	err, user := proxy.GetUserAttribute(userId)
	if err != nil {
		return nil, err
	}
	resp := &cgi.LoginRsp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		CodeSession: session,
		User:        user,
	}
	return resp, err
}

// 自己桌球登陆
func (h *cgiProxyService) ByteDanceBilliardLogin(ctx context.Context, in *cgi.LoginReq) (*cgi.LoginRsp, error) {
	log.Printf("billiard bd login %v", in)
	err, session := platform.ByteDanceCode2Session(in.Code, ByteDanceBilliardAppId, ByteDanceBilliardAppSecret)
	if err != nil {
		return nil, err
	}
	gameName := "bytedance_billiard"
	err, userId := proxy.LoginGetUserId(session, ByteDanceBilliardAppId, gameName)
	if err != nil {
		return nil, err
	}
	err, user := proxy.GetUserAttribute(userId)
	if err != nil {
		return nil, err
	}
	resp := &cgi.LoginRsp{
		Header: &cgi.Header{
			Code:    0,
			Message: "ok",
		},
		CodeSession: session,
		User:        user,
	}
	return resp, err
}
