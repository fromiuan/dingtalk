package dingtalk

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/fromiuan/dingtalk/lib"
)

type AuthClient struct {
	AppId     string
	AppSecret string
}

// 扫描登陆第三方、钉钉内部登录第三方、密码登录第三方
type UserinfoBycode struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	UserInfo struct {
		Nick    string `json:"nick"`
		Openid  string `json:"openid"`
		Unionid string `json:"unionid"`
	} `json:"user_info"`
}

// 企业内部应用免登
type UserinfoByInternal struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	UserID   string `json:"userid"`
	SYSLevel int    `json:"sys_level"`
	IsSYS    bool   `json:"is_sys"`
}

// 应用管理后台免登
type SsoAccessToken struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
}
type UserinfoBySso struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	IsSys    bool   `json:"is_sys"`
	UserInfo struct {
		Avatar string `json:"avatar"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		UserID string `json:"userid"`
	} `json:"user_info"`
	CorpInfo struct {
		CorpName string `json:"corp_name"`
		CorpID   string `json:"corpid"`
	} `json:"corp_info"`
}

// 扫描登陆第三方、钉钉内部登录第三方、密码登录第三方
func (c *Client) GetUserInfoByCode(code, appId, appSecret string) (u *UserinfoBycode, err error) {
	u = new(UserinfoBycode)
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	paramData := map[string]string{
		"tmp_auth_code": code,
	}
	jsonData, err := json.Marshal(paramData)
	if err != nil {
		return u, err
	}
	sign := lib.UserSignatur(timestamp, appSecret)
	param := map[string]string{
		"accessKey": appId,
		"timestamp": timestamp,
		"signature": sign,
	}
	url := fmt.Sprintf("%s?%s", GetUserInfoBycode, lib.UrlEncode(param))
	req := lib.Post(url)
	req.Body(jsonData)
	b, err := req.AsBytes()
	if err != nil {
		return u, err
	}

	err = json.Unmarshal(b, u)
	if err != nil {
		return u, err
	}
	return u, nil
}

// 企业内部应用免登
func (c *Client) GetUserInfoByIn(code string) (u *UserinfoByInternal, err error) {
	u = new(UserinfoByInternal)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return u, err
	}

	url := fmt.Sprintf("%s?access_token=%s", GetUserInfo, accessToken)
	if c.Debug {
		log.Println("url:", url)
	}
	var body []byte
	body, err = lib.Get(url).AsBytes()
	if err != nil {
		return u, err
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		return u, err
	}
	return u, nil
}

// 应用管理后台免登
func (c *Client) GetUserInfoBySso(code, ssoAccessToken string) (*UserinfoBySso, error) {
	u := new(UserinfoBySso)
	if ssoAccessToken == "" {
		accessTokenInfo, err := c.GetSsoAccessToken(code)
		if err != nil {
			return u, err
		}
		ssoAccessToken = accessTokenInfo.AccessToken
	}

	url := fmt.Sprintf("%s?access_token=%s&code=%s", SSOGetUserInfo, ssoAccessToken, code)
	if c.Debug {
		log.Println("url:", url)
	}
	var body []byte
	body, err := lib.Get(url).AsBytes()
	if err != nil {
		return u, err
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		return u, err
	}
	return u, nil
}

// 应用管理后台免登获取AccessToken
func (c *Client) GetSsoAccessToken(code string) (*SsoAccessToken, error) {
	s := new(SsoAccessToken)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return s, err
	}

	url := fmt.Sprintf("%s?access_token=%s&code=%s", SSOGetToken, accessToken, code)
	if c.Debug {
		log.Println("url:", url)
	}
	var body []byte
	body, err = lib.Get(url).AsBytes()
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(body, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}
