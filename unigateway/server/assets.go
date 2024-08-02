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

func daoAssetToCgi(at *pb.GameAsset) *cgi.GameAsset {
	return &cgi.GameAsset{
		Id:        at.Id,
		Name:      at.Name,      // 资产名称
		Type:      at.Type,      // 资产类型
		Level:     at.Level,     // 资产级别
		GoldValue: at.GoldValue, // 资产价值
		Url:       at.Url,       // 资产图地址
	}
}

func daoUserAssetToCgi(at *pb.UserAsset) *cgi.UserAsset {
	return &cgi.UserAsset{
		Id:             at.Id,
		UserId:         at.UserId,               // 用户ID
		AssetId:        at.AssetId,              // 资产ID
		Count:          at.Count,                // 拥有的数量
		ExpirationTime: at.ExpirationTime,       // 过期时间
		Asset:          daoAssetToCgi(at.Asset), // 资产项内容
	}
}

func AssetList(userId uint64) (error, []*cgi.GameAsset) {
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

	req := &pb.CommonUserReq{
		UserId: userId,
	}
	r, err := c.AssetList(ctx, req)
	if err != nil {
		log.Fatalf("asset list failed: %v", err)
		return err, nil
	}

	assetList := make([]*cgi.GameAsset, len(r.Assets))
	for i := 0; i < len(r.Assets); i++ {
		assetList[i] = daoAssetToCgi(r.Assets[i])
	}

	return nil, assetList
}

// 查询用户的资产列表
func UserAssetList(userId uint64) (error, []*cgi.UserAsset) {
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

	req := &pb.CommonUserReq{
		UserId: userId,
	}
	r, err := c.UserAssetList(ctx, req)
	if err != nil {
		log.Fatalf("asset list failed: %v", err)
		return err, nil
	}

	userAssetList := make([]*cgi.UserAsset, len(r.UserAssets))
	for i := 0; i < len(r.UserAssets); i++ {
		userAssetList[i] = daoUserAssetToCgi(r.UserAssets[i])
	}

	return nil, userAssetList
}

// 查询用户的资产列表
func UpdateUserAsset(userId uint64, assetId uint64, expirationTime string) (error, []*cgi.UserAsset) {
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

	req := &pb.UpdateUserAssetReq{
		UserId:         userId,
		AssetId:        assetId,
		ExpirationTime: expirationTime,
	}
	r, err := c.UpdateUserAsset(ctx, req)
	if err != nil {
		log.Fatalf("asset list failed: %v", err)
		return err, nil
	}

	userAssetList := make([]*cgi.UserAsset, len(r.UserAssets))
	for i := 0; i < len(r.UserAssets); i++ {
		userAssetList[i] = daoUserAssetToCgi(r.UserAssets[i])
	}

	return nil, userAssetList
}
