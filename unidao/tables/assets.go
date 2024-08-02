package tables

import (
	"errors"
	"log"

	pb "github.com/qwenyang/xmetau/proto/unidao"
)

type TableGameAsset struct {
	ID        uint64 `gorm:"column:FId"`
	Name      string `gorm:"column:FName"`
	Type      int32  `gorm:"column:FType"`
	Level     int32  `gorm:"column:FLevel"`
	GoldValue int32  `gorm:"column:FGoldValue"`
	Url       string `gorm:"column:FUrl"`
}

type TableUserAsset struct {
	ID             uint64         `gorm:"column:FId"`
	UserId         uint64         `gorm:"column:FUserId"`
	AssetId        uint64         `gorm:"column:FAssetId"`
	Count          int32          `gorm:"column:FCount"`
	ExpirationTime string         `gorm:"column:FExpirationTime"`
	GameAsset      TableGameAsset `gorm:"foreignKey:ID;references:AssetId"`
}

func (tu *TableGameAsset) TableName() string {
	return "T_GameAssets"
}

func (tu *TableUserAsset) TableName() string {
	return "T_UserAssets"
}

func newUserAsset(req *pb.UpdateUserAssetReq) *TableUserAsset {
	ua := &TableUserAsset{
		ID:             0,
		UserId:         req.UserId,
		AssetId:        req.AssetId,
		Count:          1,
		ExpirationTime: req.ExpirationTime,
	}
	return ua
}

func tableAssetToProto(at *TableGameAsset) *pb.GameAsset {
	return &pb.GameAsset{
		Id:        at.ID,
		Name:      at.Name,      // 资产名称
		Type:      at.Type,      // 资产类型
		Level:     at.Level,     // 资产级别
		GoldValue: at.GoldValue, // 资产价值
		Url:       at.Url,       // 资产图地址
	}
}

func tableUserAssetToProto(at *TableUserAsset) *pb.UserAsset {
	return &pb.UserAsset{
		Id:             at.ID,
		UserId:         at.UserId,
		AssetId:        at.AssetId,
		Count:          at.Count,
		ExpirationTime: at.ExpirationTime,
		Asset:          tableAssetToProto(&at.GameAsset),
	}
}

// 查询资产列表
func AssetList(user *pb.CommonUserReq) (error, []*pb.GameAsset) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	assets := make([]TableGameAsset, 1001)
	db.Limit(1000).Find(&assets)

	assetList := make([]*pb.GameAsset, len(assets))
	for i := 0; i < len(assets); i++ {
		assetList[i] = tableAssetToProto(&assets[i])
	}

	return nil, assetList
}

// 查询用户的资产列表
func UserAssetList(user *pb.CommonUserReq) (error, []*pb.UserAsset) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	userAssets := make([]TableUserAsset, 101)
	log.Printf("UserAssetList %v", user)
	db.Limit(100).Model(&TableUserAsset{}).Preload("GameAsset").Where("FUserId = ? ", user.UserId).Find(&userAssets)

	userAssetList := make([]*pb.UserAsset, len(userAssets))
	for i := 0; i < len(userAssets); i++ {
		userAssetList[i] = tableUserAssetToProto(&userAssets[i])
	}

	return nil, userAssetList
}

// 更新用户资产信息
func UpdateUserAsset(at *pb.UpdateUserAssetReq) (error, []*pb.UserAsset) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}

	uat := newUserAsset(at)
	db.Create(uat)

	req := &pb.CommonUserReq{
		UserId: at.UserId,
	}
	return UserAssetList(req)
}
