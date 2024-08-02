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

type WxCode2SessionResp struct {
	Errcode    int    `json:"errcode,omitempty"`
	Errmsg     string `json:"errmsg,omitempty"`
	Unionid    string `json:"unionid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	Openid     string `json:"openid,omitempty"`
}

func LoginGetUserId(in *cgi.CodeSession, appID, gameName string) (error, uint64) {
	log.Printf("login get userid: %v", in)
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewXMetauDaoServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if in.Unionid == "" {
		in.Unionid = in.Openid + "_" + appID
	}
	loginReq := &pb.LoginReq{
		Login: &pb.Login{
			LoginType: 0,
			AppId:     appID,
			OpenId:    in.Openid,
			UnionId:   in.Unionid,
		},
		GameName: gameName,
	}

	resp, err := c.Login(ctx, loginReq)
	if err != nil {
		log.Fatalf("login failed: %v", err)
		return err, 0
	}

	log.Printf("login resp: %v", resp)

	return nil, resp.Account.UserId
}
