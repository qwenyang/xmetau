package tables

import (
	"errors"
	"log"
	"time"

	pb "github.com/qwenyang/xmetau/proto/unidao"
	"gorm.io/gorm"
)

type LoginAccount struct {
	UserId    uint64 `gorm:"column:FUserId"`
	LoginType int32  `gorm:"column:FLoginType"`
	AppId     string `gorm:"column:FAppId"`
	OpenId    string `gorm:"column:FOpenId"`
	UnionId   string `gorm:"column:FUnionId"`
}

func (tu *LoginAccount) TableName() string {
	return "T_LoginAccount"
}

func newLoginAccount(reqLogin *pb.Login) *LoginAccount {
	account := LoginAccount{
		UserId:    uint64(time.Now().UnixMicro()),
		LoginType: reqLogin.LoginType,
		AppId:     reqLogin.AppId,
		OpenId:    reqLogin.OpenId,
		UnionId:   reqLogin.UnionId,
	}
	return &account
}

func newLoginByAccount(acount *LoginAccount) *pb.Account {
	account := &pb.Account{
		UserId: acount.UserId,
	}
	return account
}

// 1. 登录
func LoginRegister(reqLogin *pb.Login) (error, bool, *pb.Account) {
	// db, err := gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{})
	db, err := openDB()
	if err != nil {
		log.Printf("failed to open mysql %v", err)
		return errors.New("db open failed"), false, nil
	}
	var account = &LoginAccount{}
	var result = db.First(account, "FOpenId = ? and FAppId = ?", reqLogin.OpenId, reqLogin.AppId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		account = newLoginAccount(reqLogin)
		db.Create(account)
		return nil, true, newLoginByAccount(account)
	}
	return nil, false, newLoginByAccount(account)
}
