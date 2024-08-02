package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/qwenyang/xmetau/unigateway/proto/cgi"
)

type ByteDanceCode2SessionResp struct {
	Errcode    int    `json:"errcode,omitempty"`
	Errmsg     string `json:"errmsg,omitempty"`
	Unionid    string `json:"unionid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	Openid     string `json:"openid,omitempty"`
}

func ByteDanceCode2Session(code string, appID string, appSecret string) (error, *cgi.CodeSession) {
	strUrl := fmt.Sprintf("https://minigame.zijieapi.com/mgplatform/api/apps/jscode2session?appid=%s&secret=%s&code=%s&grant_type=authorization_code", appID, appSecret, code)
	client := &http.Client{}
	req, err := http.NewRequest("GET", strUrl, nil)
	if err != nil {
		log.Println("HttpRequestFail", code, err)
		return errors.New("ErrSysHttpFail"), nil
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("HttpRequestDoFail", code, err)
		return errors.New("ErrSysHttpFail"), nil
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		log.Println("HttpReadAllFail", code, err)
		return errors.New("ErrSysHttpFail"), nil
	}
	bdRsp := &ByteDanceCode2SessionResp{}
	json.Unmarshal(bodyText, bdRsp)
	log.Println("BdCode2Session", code, *bdRsp, string(bodyText))
	if bdRsp.Errcode != 0 {
		log.Println("BdCode2SessionFail", code, *bdRsp, string(bodyText))
		return errors.New("ErrIvalidMiniCode"), nil
	}

	pbRsp := &cgi.CodeSession{
		Openid:     bdRsp.Openid,
		SessionKey: bdRsp.SessionKey,
		Unionid:    bdRsp.Unionid,
	}
	return nil, pbRsp
}
