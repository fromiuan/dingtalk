package dingtalk

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fromiuan/dingtalk/lib"
)

// 创建、更新用户
type CreateUser struct {
	UserID       string      `json:"userid"`
	Name         string      `json:"name"`
	OrderInDepts string      `json:"orderInDepts"`
	Department   []int       `json:"department"`
	Position     string      `json:"position"`
	Mobile       string      `json:"mobile"`
	Tel          string      `json:"tel"`
	WorkPlace    string      `json:"workPlace"`
	Remark       string      `json:"remark"`
	Email        string      `json:"email"`
	OrgEmail     string      `json:"orgEmail"`
	Jobnumber    string      `json:"jobnumber"`
	IsHide       bool        `json:"isHide"`
	IsSenior     bool        `json:"isSenior"`
	Extattr      interface{} `json:"extattr"`
}

type CreateUserRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	UserID  string `json:"userid"`
}

// 获取部门用户userid列表
type UserGetDeptMemberRsp struct {
	Errcode int      `json:"errcode"`
	Errmsg  string   `json:"errmsg"`
	UserIds []string `json:"userIds"`
}

// 获取用户详情
type UsertGetRsp struct {
	Errcode         int    `json:"errcode"`
	UnionID         string `json:"unionid"`
	Remark          string `json:"remark"`
	UserID          string `json:"userid"`
	IsLeaderInDepts string `json:"isLeaderInDepts"`
	IsBoss          bool   `json:"isBoss"`
	HiredDate       int64  `json:"hiredDate"`
	IsSenior        bool   `json:"isSenior"`
	Tel             string `json:"tel"`
	Department      []int  `json:"department"`
	WorkPlace       string `json:"workPlace"`
	Email           string `json:"email"`
	OrderInDepts    string `json:"orderInDepts"`
	Mobile          string `json:"mobile"`
	Errmsg          string `json:"errmsg"`
	Active          bool   `json:"active"`
	Avatar          string `json:"avatar"`
	IsAdmin         bool   `json:"isAdmin"`
	IsHide          bool   `json:"isHide"`
	Jobnumber       string `json:"jobnumber"`
	Name            string `json:"name"`
	Extattr         struct {
	} `json:"extattr"`
	StateCode string `json:"stateCode"`
	Position  string `json:"position"`
	Roles     []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		GroupName string `json:"groupName"`
	} `json:"roles"`
}

// 创建用户
func (c *Client) UserCreate(cu *CreateUser) (*CreateUserRsp, error) {
	rsp := new(CreateUserRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	jsonParms, err := json.Marshal(cu)
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s", UserCreate, accessToken)
	if c.Debug {
		log.Println("url:", url)
	}
	body, err := lib.Post(url).Body(jsonParms).AsBytes()
	if err != nil {
		return rsp, err
	}
	err = json.Unmarshal(body, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// 更新用户
func (c *Client) UserUpdate(cu *CreateUser) (*CreateUserRsp, error) {
	rsp := new(CreateUserRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	jsonParms, err := json.Marshal(cu)
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s", UserUpdate, accessToken)
	if c.Debug {
		log.Println("url:", url)
	}
	body, err := lib.Post(url).Body(jsonParms).AsBytes()
	if err != nil {
		return rsp, err
	}
	err = json.Unmarshal(body, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// 删除用户
func (c *Client) UserDelete(userid string) (*CreateUserRsp, error) {
	rsp := new(CreateUserRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s&userid=%s", UserUpdate, accessToken, userid)
	if c.Debug {
		log.Println("url:", url)
	}
	body, err := lib.Get(url).AsBytes()
	if err != nil {
		return rsp, err
	}
	err = json.Unmarshal(body, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// 获取用户详情
func (c *Client) UserGet(userid string) (*UsertGetRsp, error) {
	rsp := new(UsertGetRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s&userid=%s", UserGet, accessToken, userid)
	if c.Debug {
		log.Println("url:", url)
	}
	body, err := lib.Get(url).AsBytes()
	if err != nil {
		return rsp, err
	}
	err = json.Unmarshal(body, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// 获取部门用户userid列表
func (c *Client) UserGetDeptMember(userid string) (*UserGetDeptMemberRsp, error) {
	rsp := new(UserGetDeptMemberRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s&userid=%s", UserGet, accessToken, userid)
	if c.Debug {
		log.Println("url:", url)
	}
	body, err := lib.Get(url).AsBytes()
	if err != nil {
		return rsp, err
	}
	err = json.Unmarshal(body, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}
