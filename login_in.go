// 企业内部登录
package dingtalk

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/fromiuan/dingtalk/lib"
)

type UserinfoByInternal struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	UserID   string `json:"userid"`
	SYSLevel int    `json:"sys_level"`
	IsSYS    bool   `json:"is_sys"`
}

func (c *Client) GetUserInfoByCode(code string) (u *UserinfoByInternal, err error) {
	u = new(UserinfoByInternal)
	access_token, err := c.GetAccessToken()
	if err != nil {
		return u, err
	}

	url := fmt.Sprintf("%s?access_token=%s", getuserinfo, access_token)

	var body []byte
	body, err = lib.Get(url).AsBytes()
	if err != nil {
		return u, err
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		return u, err
	}
	if u.ErrCode != 0 {
		return u, errors.New(u.ErrMsg)
	}
	return u, nil
}
