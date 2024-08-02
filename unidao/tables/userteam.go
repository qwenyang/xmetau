package tables

import (
	"errors"
	"log"

	pb "github.com/qwenyang/xmetau/proto/unidao"
	"gorm.io/gorm"
)

type TableTeamMember struct {
	TeamId        uint64             `gorm:"column:FTeamId"`
	UserId        uint64             `gorm:"column:FUserId"`
	CharType      int32              `gorm:"column:FCharType"`
	LevelIndex    int32              `gorm:"column:FLevelIndex"`
	PassCount     int32              `gorm:"column:FPassCount"`
	FinishStatus  int32              `gorm:"column:FFinishStatus"`
	Season        int32              `gorm:"column:FSeason"`
	UserAttribute TableUserAttribute `gorm:"foreignKey:UserId;references:UserId"`
}

type TableTeamGroup struct {
	TeamId      uint64            `gorm:"column:FTeamId"`
	TeamName    string            `gorm:"column:FTeamName"`
	LevelIndex  int32             `gorm:"column:FLevelIndex"`
	Score       int32             `gorm:"column:FScore"`
	Season      int32             `gorm:"column:FSeason"`
	TeamMembers []TableTeamMember `gorm:"foreignKey:TeamId;references:TeamId"`
}

type WriteTableTeamGroup struct {
	TeamId     uint64 `gorm:"column:FTeamId"`
	TeamName   string `gorm:"column:FTeamName"`
	LevelIndex int32  `gorm:"column:FLevelIndex"`
	Score      int32  `gorm:"column:FScore"`
	Season     int32  `gorm:"column:FSeason"`
}

func (tu *TableTeamGroup) TableName() string {
	return "T_Team"
}
func (tu *WriteTableTeamGroup) TableName() string {
	return "T_Team"
}
func (tu *TableTeamMember) TableName() string {
	return "T_TeamMember"
}

func newWriteTableTeamGroup(teamId uint64, teamName string) *WriteTableTeamGroup {
	tmm := &WriteTableTeamGroup{
		TeamId:     teamId,
		TeamName:   teamName,
		LevelIndex: 0,
		Score:      0,
		Season:     0,
	}
	return tmm
}

func newTableTeamMember(userId uint64, teamId uint64, charType int32) *TableTeamMember {
	tmm := &TableTeamMember{
		TeamId:       teamId,
		UserId:       userId,
		CharType:     charType,
		LevelIndex:   0,
		PassCount:    0,
		FinishStatus: 0,
		Season:       0,
	}
	return tmm
}

func tableTeamGroupToProto(ttg *TableTeamGroup) *pb.TeamGroup {
	tg := &pb.TeamGroup{
		TeamId:      ttg.TeamId,
		TeamName:    ttg.TeamName,
		LevelIndex:  ttg.LevelIndex,
		Score:       ttg.Score,
		TeamMembers: make([]*pb.TeamMember, len(ttg.TeamMembers)),
	}
	for i := 0; i < len(ttg.TeamMembers); i++ {
		tg.TeamMembers[i] = &pb.TeamMember{
			TeamId:       ttg.TeamMembers[i].TeamId,
			UserId:       ttg.TeamMembers[i].UserId,
			CharType:     ttg.TeamMembers[i].CharType,
			LevelIndex:   ttg.TeamMembers[i].LevelIndex,
			PassCount:    ttg.TeamMembers[i].PassCount,
			FinishStatus: ttg.TeamMembers[i].FinishStatus,
			UserAttribute: &pb.UserAttribute{
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

// 获取用户组队信息
func GetUserTeam(userId uint64) (error, *pb.TeamGroup) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	ut := newTableTeamMember(userId, 0, 0)
	var result = db.First(ut, "FUserId = ? and FSeason = 0", userId)
	log.Printf("user team first %v ", result)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Printf("user userid=%v is not found", userId)
		return nil, nil
	}
	// 通过TeamId来加载各种信息
	tg := new(TableTeamGroup)
	db.Model(&TableTeamGroup{}).Preload("TeamMembers").Preload("TeamMembers.UserAttribute").Where("FTeamId = ? and FSeason = 0", ut.TeamId).First(&tg)

	team := tableTeamGroupToProto(tg)

	return nil, team
}

// 获取用户组队信息
func GetTeamInfo(teamId uint64) (error, *pb.TeamGroup) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	// 通过TeamId来加载各种信息
	tg := new(TableTeamGroup)
	db.Model(&TableTeamGroup{}).Preload("TeamMembers").Preload("TeamMembers.UserAttribute").Where("FTeamId = ? and FSeason = 0", teamId).First(&tg)

	team := tableTeamGroupToProto(tg)

	return nil, team
}

// 更新用户资产信息
func CreateTeam(req *pb.CreateTeamReq) (error, *pb.TeamGroup) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}

	wtt := newWriteTableTeamGroup(req.TeamId, req.TeamName)
	wtt.LevelIndex = 1
	db.Create(wtt)

	// 本来应该放到一个事务里面处理比较好，这里即使Team成功，用户失败，可以换个teamID重新创建也OK
	ttm := newTableTeamMember(req.UserId, req.TeamId, req.CharType)
	ttm.LevelIndex = 1
	db.Create(ttm)
	return GetUserTeam(req.UserId)
}

// 加入新用户
func JoinTeam(req *pb.JoinTeamReq) (error, *pb.TeamGroup) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}

	ttm := newTableTeamMember(req.UserId, req.TeamId, req.CharType)
	db.Create(ttm)
	return GetUserTeam(req.UserId)
}

// 队伍排名类别
func GetTeamRankList(req *pb.TeamRankListReq) (error, []*pb.TeamGroup) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	teams := make([]TableTeamGroup, 101)
	// 通过当前的棋力值 + 差值计算最终的棋力
	db.Limit(100).Order("FScore desc").Where("FSeason = 0").Find(&teams)

	tgs := make([]*pb.TeamGroup, len(teams))
	for i := 0; i < len(teams); i++ {
		tgs[i] = tableTeamGroupToProto(&teams[i])
	}

	return nil, tgs
}

// 加入新用户
func UpdateUserTeam(req *pb.UpdateUserTeamReq) (error, *pb.TeamGroup) {
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), nil
	}
	tm := newTableTeamMember(req.UserId, req.TeamId, 0)
	if req.UpdateType == 1 { // 1 开始游戏
		// 增加一次闯关数量
		db.Model(tm).Where("FUserId = ? and FTeamId = ? and FSeason = 0", req.UserId, req.TeamId).Updates(map[string]interface{}{
			"FPassCount": req.NewUserResource.PassCount,
		})
		return GetUserTeam(req.UserId)
	}
	if req.UpdateType == 2 { // 2 游戏通关
		// 更新用户信息
		db.Model(tm).Where("FUserId = ? and FTeamId = ? and FSeason = 0", req.UserId, req.TeamId).Updates(map[string]interface{}{
			"FPassCount":    req.NewUserResource.PassCount,
			"FLevelIndex":   req.NewUserResource.LevelIndex,
			"FFinishStatus": req.NewUserResource.FinishStatus,
		})
		// 更新团队信息
		tmg := newWriteTableTeamGroup(req.TeamId, "")
		db.Model(tmg).Where("FTeamId = ? and FLevelIndex = ?", req.TeamId, req.OldTeamResource.LevelIndex).Updates(map[string]interface{}{
			"FLevelIndex": req.NewTeamResource.LevelIndex,
			"FScore":      req.NewTeamResource.Score,
		})
		return GetUserTeam(req.UserId)
	}
	if req.UpdateType == 4 { // 4 同步用户数据
		// 增加一次闯关数量
		db.Model(tm).Where("FUserId = ? and FTeamId = ? and FSeason = 0", req.UserId, req.TeamId).Updates(map[string]interface{}{
			"FPassCount":    req.NewUserResource.PassCount,
			"FLevelIndex":   req.NewUserResource.LevelIndex,
			"FFinishStatus": req.NewUserResource.FinishStatus,
		})
		return GetUserTeam(req.UserId)
	}
	if req.UpdateType == 5 { // 5 更新结束标识
		db.Model(tm).Where("FUserId = ? and FTeamId = ? and FSeason = 0", req.UserId, req.TeamId).Updates(map[string]interface{}{
			"FFinishStatus": req.NewUserResource.FinishStatus,
		})
		return GetUserTeam(req.UserId)
	}
	if req.UpdateType == 10 { // 10 解散团队
		// 解散团队成员
		db.Model(tm).Where("FTeamId = ? and FSeason = 0", req.TeamId).Updates(map[string]interface{}{
			"FSeason": 1,
		})
		// 解散团队
		tmg := newWriteTableTeamGroup(req.TeamId, "")
		db.Model(tmg).Where("FTeamId = ? and FSeason = 0", req.TeamId).Updates(map[string]interface{}{
			"FSeason": 1,
		})
		return GetUserTeam(req.UserId)
	}
	return GetUserTeam(req.UserId)
}
