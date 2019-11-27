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

// 获取部门用户
type UserSimpleListRsp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	HasMore  bool   `json:"hasMore"`
	Userlist []struct {
		Userid string `json:"userid"`
		Name   string `json:"name"`
	} `json:"userlist"`
}

// 未登录钉钉的员工列表
type UserGetInactiveRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  struct {
		HasMore bool     `json:"has_more"`
		List    []string `json:"list"`
	} `json:"result"`
}

// 获取部门用户详情
type UserListByPageRsp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	HasMore  bool   `json:"hasMore"`
	Userlist []struct {
		Userid     string      `json:"userid"`
		Unionid    string      `json:"unionid"`
		Mobile     string      `json:"mobile"`
		Tel        string      `json:"tel"`
		WorkPlace  string      `json:"workPlace"`
		Remark     string      `json:"remark"`
		Order      int         `json:"order"`
		IsAdmin    bool        `json:"isAdmin"`
		IsBoss     bool        `json:"isBoss"`
		IsHide     bool        `json:"isHide"`
		IsLeader   bool        `json:"isLeader"`
		Name       string      `json:"name"`
		Active     bool        `json:"active"`
		Department []int       `json:"department"`
		Position   string      `json:"position"`
		Email      string      `json:"email"`
		Avatar     string      `json:"avatar"`
		Jobnumber  string      `json:"jobnumber"`
		Extattr    interface{} `json:"extattr"`
	} `json:"userlist"`
}

// 获取管理员列表
type UserGetAdminRsp struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	AdminList []struct {
		SysLevel int    `json:"sys_level"`
		Userid   string `json:"userid"`
	} `json:"admin_list"`
}

// 获取管理员通讯录权限范围
type UserGetAdminScopeRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	DeptIds []int  `json:"dept_ids"`
}

// 根据unionid获取userid
type UserGetUseridByUnionidRsp struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	ContactType int    `json:"contactType"`
	Userid      string `json:"userid"`
}

// 根据手机号获取userid
type UserGetByMobileRsp struct {
	CreateUserRsp
}

// 获取企业员工人数
type UserGetOrgUserCountRsp struct {
	Count   int    `json:"count"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// 创建用户
func (c *Client) UserCreate(user *CreateUser) (*CreateUserRsp, error) {
	rsp := new(CreateUserRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	jsonParms, err := json.Marshal(user)
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
func (c *Client) UserUpdate(user *CreateUser) (*CreateUserRsp, error) {
	rsp := new(CreateUserRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	jsonParms, err := json.Marshal(user)
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

	url := fmt.Sprintf("%s?access_token=%s", UserGetDeptMember, accessToken)
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

// 获取部门用户
func (c *Client) UserSimpleList(departmentid int) (*UserSimpleListRsp, error) {
	rsp := new(UserSimpleListRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s&department_id=%d", UserSimpleList, accessToken, departmentid)
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

// 获取部门用户详情
func (c *Client) UserListByPage(departmentid int) (*UserListByPageRsp, error) {
	rsp := new(UserListByPageRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s&department_id=%d", UserListByPage, accessToken, departmentid)
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

// 获取管理员列表
func (c *Client) UserGetAdmin() (*UserGetAdminRsp, error) {
	rsp := new(UserGetAdminRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s", UserGetAdmin, accessToken)
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

// 获取管理员通讯录权限范围
func (c *Client) UserGetAdminScope() (*UserGetAdminScopeRsp, error) {
	rsp := new(UserGetAdminScopeRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s", UserGetAdminScope, accessToken)
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

// 根据unionid获取userid
func (c *Client) UserGetUseridByUnionid() (*UserGetUseridByUnionidRsp, error) {
	rsp := new(UserGetUseridByUnionidRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s", UserGetUseridByUnionid, accessToken)
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

// 根据手机号获取userid
func (c *Client) UserGetByMobile() (*UserGetByMobileRsp, error) {
	rsp := new(UserGetByMobileRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s", UserGetByMobile, accessToken)
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

// 获取企业员工人数
func (c *Client) UserGetOrgUserCount() (*UserGetOrgUserCountRsp, error) {
	rsp := new(UserGetOrgUserCountRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s", UserGetOrgUserCount, accessToken)
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

// 未登录钉钉的员工列表
func (c *Client) UserGetInactive() (*UserGetInactiveRsp, error) {
	rsp := new(UserGetInactiveRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s", UserGetInactive, accessToken)
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
