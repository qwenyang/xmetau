package tables

import (
	"errors"
	"log"

	pb "github.com/qwenyang/xmetau/proto/unidao"
	"gorm.io/gorm/clause"
)

type TableGameSetting struct {
	SetId    uint64 `gorm:"column:FSetId"`
	SetType  int32  `gorm:"column:FSetType"`
	SetKey   string `gorm:"column:FSetKey"`
	SetValue string `gorm:"column:FSetValue"`
}

func (tu *TableGameSetting) TableName() string {
	return "T_GameSetting"
}

func newGameSetting(req *pb.UpdateSettingReq) *TableGameSetting {
	ua := &TableGameSetting{
		SetId:    req.SetId,
		SetType:  req.SetType,
		SetKey:   req.SetKey,
		SetValue: req.SetValue,
	}
	return ua
}

func tableSettingToProto(gs *TableGameSetting) *pb.GameSetting {
	return &pb.GameSetting{
		SetId:    gs.SetId,
		SetType:  gs.SetType,
		SetKey:   gs.SetKey,
		SetValue: gs.SetValue,
	}
}

// 设置列表
func SettingList(user *pb.GameSettingReq) (error, []*pb.GameSetting) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	settings := make([]TableGameSetting, 101)
	log.Printf("settings %v", settings)
	db.Limit(100).Model(&TableGameSetting{}).Where("FSetId = ?", user.SetId).Find(&settings)

	settingsList := make([]*pb.GameSetting, len(settings))
	for i := 0; i < len(settings); i++ {
		settingsList[i] = tableSettingToProto(&settings[i])
	}

	return nil, settingsList
}

// 更新用户资产信息
func UpdateSetting(at *pb.UpdateSettingReq) (error, []*pb.GameSetting) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}

	uat := newGameSetting(at)
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "FSetId"}, {Name: "FSetType"}, {Name: "FSetKey"}},
		DoUpdates: clause.AssignmentColumns([]string{"FSetValue"}),
	}).Create(uat)

	req := &pb.GameSettingReq{
		UserId:  at.UserId,
		SetId:   at.SetId,
		SetType: at.SetType,
	}
	return SettingList(req)
}
