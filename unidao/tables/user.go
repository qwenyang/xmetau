package tables

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"gorm.io/gorm"

	pb "github.com/qwenyang/xmetau/proto/unidao"
)

type TableUserAttribute struct {
	UserId         uint64 `gorm:"column:FUserId"`
	NickName       string `gorm:"column:FNickName"`
	AvatarUrl      string `gorm:"column:FAvatarUrl"`
	NoviceTraining int32  `gorm:"column:FNoviceTraining"`
	PlayLevel      int32  `gorm:"column:FPlayLevel"`
	GoldCoin       int32  `gorm:"column:FGoldCoin"`
	WinNum         int32  `gorm:"column:FWinNum"`
	LoseNum        int32  `gorm:"column:FLoseNum"`
	TieNum         int32  `gorm:"column:FTieNum"`
	ModifyTime     string `gorm:"column:FModifyTime"`
	CreateTime     string `gorm:"column:FCreateTime"`
}

type WriteTableUserAttribute struct {
	UserId         uint64 `gorm:"column:FUserId"`
	NickName       string `gorm:"column:FNickName"`
	AvatarUrl      string `gorm:"column:FAvatarUrl"`
	NoviceTraining int32  `gorm:"column:FNoviceTraining"`
	PlayLevel      int32  `gorm:"column:FPlayLevel"`
	GoldCoin       int32  `gorm:"column:FGoldCoin"`
	WinNum         int32  `gorm:"column:FWinNum"`
	LoseNum        int32  `gorm:"column:FLoseNum"`
	TieNum         int32  `gorm:"column:FTieNum"`
	GameName       string `gorm:"column:FGameName"`
}

func (tu *TableUserAttribute) TableName() string {
	return "T_UserAttribute"
}

func (tu *WriteTableUserAttribute) TableName() string {
	return "T_UserAttribute"
}

func newUserAttribute(userId uint64) *TableUserAttribute {
	userInfo := &TableUserAttribute{
		UserId:         userId,
		NickName:       "",
		AvatarUrl:      "",
		NoviceTraining: 0,
		PlayLevel:      0,
		GoldCoin:       0,
		WinNum:         0,
		LoseNum:        0,
		TieNum:         0,
		ModifyTime:     "",
		CreateTime:     "",
	}
	return userInfo
}

func tableUserToProto(userInfo *TableUserAttribute) *pb.UserAttribute {
	return &pb.UserAttribute{
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

// 更新用户表信息数据
func UpdateUserTable(user *pb.UserAttribute) (error, *pb.UserAttribute) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	userInfo := &WriteTableUserAttribute{
		UserId:         user.UserId,
		NickName:       user.NickName,
		AvatarUrl:      user.AvatarUrl,
		NoviceTraining: user.NoviceTraining,
		PlayLevel:      user.PlayLevel,
		GoldCoin:       user.GoldCoin,
		WinNum:         user.WinNum,
		LoseNum:        user.LoseNum,
		TieNum:         user.TieNum,
		GameName:       user.GameName,
	}
	var result = db.First(&WriteTableUserAttribute{}, "FUserId = ?", user.UserId)
	log.Printf("user table first %v ", result)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Printf("create new user record %v", userInfo)
		db.Create(userInfo)
	} else {
		log.Printf("save a record")
		db.Model(userInfo).Where("FUserId = ?", user.UserId).Updates(userInfo)
	}
	return nil, user
}

// 查询户表信息数据
func GetUserTable(userId uint64) (error, *pb.UserAttribute) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	userInfo := newUserAttribute(userId)
	var result = db.First(userInfo, "FUserId = ?", userId)
	log.Printf("user table first %v ", result)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Printf("user userid=%v is not found", userId)

		createUserInfo := &WriteTableUserAttribute{
			UserId:         userId,
			NickName:       "",
			AvatarUrl:      "",
			NoviceTraining: 0,
			PlayLevel:      0,
			GoldCoin:       10000,
			WinNum:         0,
			LoseNum:        0,
			TieNum:         0,
		}
		db.Create(createUserInfo)

		db.First(userInfo, "FUserId = ?", userId)
	}
	return nil, tableUserToProto(userInfo)
}

// 更新用户头不信息
func UpdateUserHeader(user *pb.UpdateUserHeaderReq) (error, *pb.UserAttribute) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	userInfo := newUserAttribute(user.UserId)
	db.Model(userInfo).Where("FUserId = ?", user.UserId).Updates(map[string]interface{}{
		"FNickName": user.NickName, "FAvatarUrl": user.AvatarUrl,
	})

	return GetUserTable(user.UserId)
}

// 更新用户新手训练信息
func UpdateUserNoviceTraining(user *pb.UpdateUserTrainingReq) (error, *pb.UserAttribute) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	userInfo := newUserAttribute(user.UserId)

	db.Model(userInfo).Where("FUserId = ?", user.UserId).Updates(map[string]interface{}{
		"FNoviceTraining": user.NoviceTraining,
	})

	return GetUserTable(user.UserId)
}

// 更新用户棋力等级
func UpdateUserPlayLevel(user *pb.UpdateUserLevelReq) (error, *pb.UserAttribute) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	userInfo := newUserAttribute(user.UserId)
	// 通过当前的棋力值 + 差值计算最终的棋力值
	var level = user.PlayLevel + user.DiffLevel
	if level < 0 {
		level = 0
	}

	db.Model(userInfo).Where("FUserId = ? and FPlayLevel = ?", user.UserId, user.PlayLevel).Updates(map[string]interface{}{
		"FPlayLevel": level,
	})

	return GetUserTable(user.UserId)
}

// 更新用户金币数量
func UpdateUserCoin(user *pb.UpdateUserCoinReq) (error, *pb.UserAttribute) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	userInfo := newUserAttribute(user.UserId)
	// 通过当前的棋力值 + 差值计算最终的棋力值
	var coin = user.GoldCoin + user.DiffCoin
	if coin < 0 {
		coin = 0
	}

	db.Model(userInfo).Where("FUserId = ? and FGoldCoin = ?", user.UserId, user.GoldCoin).Updates(map[string]interface{}{
		"FGoldCoin": coin,
	})

	return GetUserTable(user.UserId)
}

// 更新游戏结束后输赢的次数
func UpdateGameNum(user *pb.UpdateGameNumReq) (error, *pb.UserAttribute) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	userInfo := newUserAttribute(user.UserId)
	// 通过当前的棋力值 + 差值计算最终的棋力值
	if user.DiffNum > 0 {
		log.Printf("win update num %v %v", user.UserId, user.WinNum)
		var win = user.WinNum + 1
		db.Model(userInfo).Where("FUserId = ? and FWinNum = ?", user.UserId, user.WinNum).Updates(map[string]interface{}{
			"FWinNum": win,
		})
	} else if user.DiffNum < 0 {
		log.Printf("Lose update num %v %v", user.UserId, user.LoseNum)
		var lose = user.LoseNum + 1
		db.Model(userInfo).Where("FUserId = ? and FLoseNum = ?", user.UserId, user.LoseNum).Updates(map[string]interface{}{
			"FLoseNum": lose,
		})
	} else {
		log.Printf("Peace update num %v %v", user.UserId, user.TieNum)
		var tie = user.TieNum + 1
		db.Model(userInfo).Where("FUserId = ? and FTieNum = ?", user.UserId, user.TieNum).Updates(map[string]interface{}{
			"FTieNum": tie,
		})
	}

	return GetUserTable(user.UserId)
}

func QueryUserRankList(user *pb.UserListReq) (error, *pb.PageData, []*pb.UserAttribute) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil, nil
	}
	users := make([]TableUserAttribute, 101)
	dfuid := 1000000000000000
	gameName := user.GetGameName()
	if user.GameName == "" {
		gameName = "douzi"
	}
	// 通过当前的棋力值 + 差值计算最终的棋力
	db.Limit(100).Order("FPlayLevel desc").Where("FNickName <> '' and FAvatarUrl <> '' and FUserId > ? and FGameName = ?", dfuid, gameName).Find(&users)

	pageData := &pb.PageData{
		Total:     100,
		PageSize:  100,
		PageIndex: 1,
	}

	userList := make([]*pb.UserAttribute, len(users))
	for i := 0; i < len(users); i++ {
		userList[i] = tableUserToProto(&users[i])
	}

	return nil, pageData, userList
}

func QueryRobotUserList(user *pb.UserListReq) (error, *pb.PageData, []*pb.UserAttribute) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil, nil
	}
	users := make([]TableUserAttribute, 1001)
	// 通过当前的棋力值 + 差值计算最终的棋力值
	db.Limit(1000).Where("FUserId < ? and FUserId > ? ORDER BY RAND()", 1000000000000000, 10000000000001).Find(&users)

	pageData := &pb.PageData{
		Total:     100,
		PageSize:  100,
		PageIndex: 1,
	}
	rand.Seed(time.Now().UnixNano()) // 纳秒时间戳
	// 随机取100个返回
	userList := make([]*pb.UserAttribute, 100)
	for i := 0; i < len(users) && i < 100; i++ {
		userList[i] = tableUserToProto(&users[i])
		pl := int32(i) * 2
		if pl > 100 {
			pl = int32(2 * rand.Int63n(50))
		}
		userList[i].PlayLevel = pl
		userList[i].GoldCoin = int32(int64(pl/10+1)*rand.Int63n(20)*1000 + 3000)
		userList[i].LoseNum = 1 + int32(rand.Int63n(10)+5*rand.Int63n(int64(pl+1)))
		userList[i].WinNum = int32(rand.Int63n(int64(userList[i].LoseNum/2+1))) + int32(userList[i].LoseNum/2)
		userList[i].TieNum = int32(rand.Int63n(3) + int64(userList[i].LoseNum/6))
	}

	return nil, pageData, userList
}
