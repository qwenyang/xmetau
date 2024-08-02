package tables

import (
	"errors"
	"log"

	pb "github.com/qwenyang/xmetau/proto/unidao"
	"gorm.io/gorm"
)

type TableUserBarrierLevel struct {
	ID         uint64 `gorm:"column:FId"`
	UserId     uint64 `gorm:"column:FUserId"`
	LevelIndex int32  `gorm:"column:FLevelIndex"`
	PassCount  int32  `gorm:"column:FPassCount"`
	PassToken  int32  `gorm:"column:FPassToken"`
}

func (tu *TableUserBarrierLevel) TableName() string {
	return "T_UserBarrierLevel"
}

func newUserBarrierLevel(req *pb.UserBarrierLevel) *TableUserBarrierLevel {
	ua := &TableUserBarrierLevel{
		ID:         0,
		UserId:     req.UserId,
		LevelIndex: req.LevelIndex,
		PassCount:  req.PassCount,
		PassToken:  req.PassToken,
	}
	return ua
}

func tableBarrierLevelToProto(at *TableUserBarrierLevel) *pb.UserBarrierLevel {
	return &pb.UserBarrierLevel{
		UserId:     at.UserId,
		LevelIndex: at.LevelIndex,
		PassCount:  at.PassCount,
		PassToken:  at.PassToken,
	}
}

// 查询用户闯关信息
func GetUserBarrierLevel(req *pb.UserBarrierLevelReq) (error, *pb.UserBarrierLevel) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	ubl := &TableUserBarrierLevel{
		ID:         0,
		UserId:     req.UserId,
		LevelIndex: 0,
		PassCount:  0,
		PassToken:  1, // 默认有一次Token机会
	}
	var result = db.First(ubl, "FUserId = ? ", req.UserId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(ubl)
		return nil, tableBarrierLevelToProto(ubl)
	}
	log.Printf("get user barrier level %v", ubl)
	return nil, tableBarrierLevelToProto(ubl)
}

// 更新用户闯关信息
func UpdateUserBarrierLevel(req *pb.UpdateUserBarrierLevelReq) (error, *pb.UserBarrierLevel) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}

	ubl := newUserBarrierLevel(req.Barrier)
	db.Model(ubl).Where("FUserId = ? and FLevelIndex <= ?", req.Barrier.UserId, req.Barrier.LevelIndex).Updates(map[string]interface{}{
		"FLevelIndex": req.Barrier.LevelIndex,
		"FPassCount":  req.Barrier.PassCount,
		"FPassToken":  req.Barrier.PassToken,
	})

	ureq := &pb.UserBarrierLevelReq{
		UserId: req.Barrier.UserId,
	}

	return GetUserBarrierLevel(ureq)
}
