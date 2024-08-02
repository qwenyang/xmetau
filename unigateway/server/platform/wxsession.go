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

type WxCode2SessionResp struct {
	Errcode    int    `json:"errcode,omitempty"`
	Errmsg     string `json:"errmsg,omitempty"`
	Unionid    string `json:"unionid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	Openid     string `json:"openid,omitempty"`
}

func WxCode2Session(code string, appID string, appSecret string) (error, *cgi.CodeSession) {
	strUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appID, appSecret, code)
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
	wxRsp := &WxCode2SessionResp{}
	json.Unmarshal(bodyText, wxRsp)
	log.Println("WxCode2Session", code, *wxRsp, string(bodyText))
	if wxRsp.Errcode != 0 {
		log.Println("WxCode2SessionFail", code, *wxRsp, string(bodyText))
		return errors.New("ErrIvalidMiniCode"), nil
	}

	pbRsp := &cgi.CodeSession{
		Openid:     wxRsp.Openid,
		SessionKey: wxRsp.SessionKey,
		Unionid:    wxRsp.Unionid,
	}
	return nil, pbRsp
}
