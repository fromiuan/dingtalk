// 企业扫描登陆
package dingtalk

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/fromiuan/dingtalk/lib"
)

type UserinfoBycode struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	UserInfo struct {
		Nick    string `json:"nick"`
		Openid  string `json:"openid"`
		Unionid string `json:"unionid"`
	} `json:"user_info"`
}

func (c *Client) GetUserInfoByCode(code string) (u *UserinfoBycode, err error) {
	u = new(UserinfoBycode)
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	paramData := map[string]string{
		"tmp_auth_code": code,
	}
	jsonData, err := json.Marshal(paramData)
	if err != nil {
		return u, err
	}
	sign := lib.UserSignatur(timestamp, c.AppSecret)
	param := map[string]string{
		"accessKey": c.AppId,
		"timestamp": timestamp,
		"signature": sign,
	}
	requestURL := getuserinfo_bycode + "?" + lib.UrlEncode(param)
	if c.Debug {
		log.Println("requestURL:", requestURL)
	}
	req := lib.Post(requestURL)
	req.Body(jsonData)
	b, err := req.AsBytes()
	if err != nil {
		return u, err
	}
	if c.Debug {
		log.Println("resp:", string(b))
	}

	err = json.Unmarshal(b, u)
	if err != nil {
		return u, err
	}
	return u, nil
}
