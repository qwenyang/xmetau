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

func daoTeamToCgi(ttg *pb.TeamGroup) *cgi.TeamGroup {
	if ttg == nil {
		return nil
	}
	tg := &cgi.TeamGroup{
		TeamId:      ttg.TeamId,
		TeamName:    ttg.TeamName,
		LevelIndex:  ttg.LevelIndex,
		Score:       ttg.Score,
		TeamMembers: make([]*cgi.TeamMember, len(ttg.TeamMembers)),
	}
	for i := 0; i < len(ttg.TeamMembers); i++ {
		tg.TeamMembers[i] = &cgi.TeamMember{
			TeamId:       ttg.TeamMembers[i].TeamId,
			UserId:       ttg.TeamMembers[i].UserId,
			CharType:     ttg.TeamMembers[i].CharType,
			LevelIndex:   ttg.TeamMembers[i].LevelIndex,
			PassCount:    ttg.TeamMembers[i].PassCount,
			FinishStatus: ttg.TeamMembers[i].FinishStatus,
			UserAttribute: &cgi.UserAttribute{
				UserId:         ttg.TeamMembers[i].UserId,
				NickName:       ttg.TeamMembers[i].UserAttribute.NickName,
				AvatarUrl:      ttg.TeamMembers[i].UserAttribute.AvatarUrl,
				NoviceTraining: ttg.TeamMembers[i].UserAttribute.NoviceTraining,
				PlayLevel:      ttg.TeamMembers[i].UserAttribute.NoviceTraining,
				GoldCoin:       ttg.TeamMembers[i].UserAttribute.GoldCoin,
				WinNum:         ttg.TeamMembers[i].UserAttribute.WinNum,
				LoseNum:        ttg.TeamMembers[i].UserAttribute.LoseNum,
				TieNum:         ttg.TeamMembers[i].UserAttribute.TieNum,
				ModifyTime:     ttg.TeamMembers[i].UserAttribute.ModifyTime,
			},
		}
	}
	return tg
}

func GetUserTeam(userId uint64) (error, *cgi.TeamGroup) {
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

	req := &pb.UserTeamReq{
		UserId: userId,
	}
	r, err := c.GetUserTeam(ctx, req)
	if err != nil {
		log.Fatalf("get user team failed: %v", err)
		return err, nil
	}

	return nil, daoTeamToCgi(r.Team)
}

func GetTeamInfo(userId uint64, teamId uint64) (error, *cgi.TeamGroup) {
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

	req := &pb.TeamInfoReq{
		UserId: userId,
		TeamId: teamId,
	}
	r, err := c.GetTeamInfo(ctx, req)
	if err != nil {
		log.Fatalf("get user team failed: %v", err)
		return err, nil
	}

	return nil, daoTeamToCgi(r.Team)
}

func CreateTeam(userId uint64, teamId uint64, teamName string, charType int32) (error, *cgi.TeamGroup) {
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

	req := &pb.CreateTeamReq{
		UserId:   userId,
		TeamId:   teamId,
		TeamName: teamName,
		CharType: charType,
	}
	r, err := c.CreateTeam(ctx, req)
	if err != nil {
		log.Fatalf("create team list failed: %v", err)
		return err, nil
	}

	return nil, daoTeamToCgi(r.Team)
}

func JoinTeam(userId uint64, teamId uint64, charType int32) (error, *cgi.TeamGroup) {
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

	req := &pb.JoinTeamReq{
		UserId:   userId,
		TeamId:   teamId,
		CharType: charType,
	}
	r, err := c.JoinTeam(ctx, req)
	if err != nil {
		log.Fatalf("create team list failed: %v", err)
		return err, nil
	}

	return nil, daoTeamToCgi(r.Team)
}

func GetTeamRankList(userId uint64) (error, []*cgi.TeamGroup) {
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

	req := &pb.TeamRankListReq{
		UserId: userId,
	}
	r, err := c.GetTeamRankList(ctx, req)
	if err != nil {
		log.Fatalf("create team list failed: %v", err)
		return err, nil
	}
	teams := make([]*cgi.TeamGroup, len(r.Teams))
	for i := 0; i < len(r.Teams); i++ {
		teams[i] = daoTeamToCgi(r.Teams[i])
	}
	return nil, teams
}

func UpdateUserTeam(req *cgi.UpdateUserTeamReq) (error, *cgi.TeamGroup) {
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

	dreq := &pb.UpdateUserTeamReq{
		UserId:     req.UserId,
		TeamId:     req.TeamId,
		UpdateType: req.UpdateType,
		OldUserResource: &pb.TeamUserResource{
			PassCount:    req.OldUserResource.PassCount,
			LevelIndex:   req.OldUserResource.LevelIndex,
			FinishStatus: req.OldUserResource.FinishStatus,
		},
		NewUserResource: &pb.TeamUserResource{
			PassCount:    req.NewUserResource.PassCount,
			LevelIndex:   req.NewUserResource.LevelIndex,
			FinishStatus: req.NewUserResource.FinishStatus,
		},
		OldTeamResource: &pb.TeamResource{
			LevelIndex: req.OldTeamResource.LevelIndex,
			Score:      req.OldTeamResource.Score,
		},
		NewTeamResource: &pb.TeamResource{
			LevelIndex: req.NewTeamResource.LevelIndex,
			Score:      req.NewTeamResource.Score,
		},
	}
	r, err := c.UpdateUserTeam(ctx, dreq)
	if err != nil {
		log.Fatalf("create team list failed: %v", err)
		return err, nil
	}

	return nil, daoTeamToCgi(r.Team)
}
