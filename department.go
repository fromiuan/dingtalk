package dingtalk

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fromiuan/dingtalk/lib"
)

// 创建部门
type CreateDepartment struct {
	Name             string `json:"name"`
	Parentid         string `json:"parentid"`
	Order            string `json:"order"`
	CreateDeptGroup  bool   `json:"createDeptGroup"`
	DeptHiding       bool   `json:"deptHiding"`
	DeptPermits      string `json:"deptPermits"`
	UserPermits      string `json:"userPermits"`
	OuterDept        bool   `json:"outerDept"`
	OuterPermitDepts string `json:"outerPermitDepts"`
	OuterPermitUsers string `json:"outerPermitUsers"`
	SourceIdentifier string `json:"sourceIdentifier"`
}

type DepartmentCreateRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	ID      int    `json:"id"`
}

// 更新部门
type UpdateDepartment struct {
	Name                  string `json:"name"`
	Parentid              string `json:"parentid"`
	Order                 string `json:"order"`
	ID                    int    `json:"id"`
	CreateDeptGroup       bool   `json:"createDeptGroup"`
	AutoAddUser           bool   `json:"autoAddUser"`
	DeptManagerUseridList string `json:"deptManagerUseridList"`
	DeptHiding            bool   `json:"deptHiding"`
	DeptPermits           string `json:"deptPermits"`
	UserPermits           string `json:"userPermits"`
	OuterDept             bool   `json:"outerDept"`
	OuterPermitDepts      string `json:"outerPermitDepts"`
	OuterPermitUsers      string `json:"outerPermitUsers"`
	OrgDeptOwner          string `json:"orgDeptOwner"`
	SourceIdentifier      string `json:"sourceIdentifier"`
}

// 获取子部门ID列表
type DepartmentListIdsRsp struct {
	Errcode       int    `json:"errcode"`
	Errmsg        string `json:"errmsg"`
	SubDeptIDList []int  `json:"sub_dept_id_list"`
}

// 获取部门列表
type DepartmentListRsp struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Department []struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		Parentid        int    `json:"parentid"`
		CreateDeptGroup bool   `json:"createDeptGroup"`
		AutoAddUser     bool   `json:"autoAddUser"`
	} `json:"department"`
}

// 获取部门详情
type DepartmentGetRsp struct {
	Errcode               int    `json:"errcode"`
	Errmsg                string `json:"errmsg"`
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	Order                 int    `json:"order"`
	Parentid              int    `json:"parentid"`
	CreateDeptGroup       bool   `json:"createDeptGroup"`
	AutoAddUser           bool   `json:"autoAddUser"`
	DeptHiding            bool   `json:"deptHiding"`
	DeptPermits           string `json:"deptPermits"`
	UserPermits           string `json:"userPermits"`
	OuterDept             bool   `json:"outerDept"`
	OuterPermitDepts      string `json:"outerPermitDepts"`
	OuterPermitUsers      string `json:"outerPermitUsers"`
	OrgDeptOwner          string `json:"orgDeptOwner"`
	DeptManagerUseridList string `json:"deptManagerUseridList"`
	SourceIdentifier      string `json:"sourceIdentifier"`
}

type DepartmentListParentRsp struct {
	Errcode    int     `json:"errcode"`
	Errmsg     string  `json:"errmsg"`
	Department [][]int `json:"department"`
}

// 创建部门
func (c *Client) DepartmentCreate(department *CreateDepartment) (*DepartmentCreateRsp, error) {
	rsp := new(DepartmentCreateRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	jsonParms, err := json.Marshal(department)
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s", DepartmentCreate, accessToken)
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

// 更新部门
func (c *Client) DepartmentUpdate(department *UpdateDepartment) (*DepartmentCreateRsp, error) {
	rsp := new(DepartmentCreateRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	jsonParms, err := json.Marshal(department)
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s", DepartmentUpdate, accessToken)
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

// 删除部门
func (c *Client) DepartmentDelete(departmentid int) (*DepartmentCreateRsp, error) {
	rsp := new(DepartmentCreateRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s&id=%d", DepartmentDelete, accessToken, departmentid)
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

// 获取子部门ID列表
func (c *Client) DepartmentListIds(departmentid int) (*DepartmentListIdsRsp, error) {
	rsp := new(DepartmentListIdsRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s&id=%d", DepartmentListIds, accessToken, departmentid)
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

// 获取部门列表
func (c *Client) DepartmentList(departmentid int) (*DepartmentListRsp, error) {
	rsp := new(DepartmentListRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s&id=%d", DepartmentList, accessToken, departmentid)
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

// 获取部门详情
func (c *Client) DepartmentGet(departmentid int) (*DepartmentGetRsp, error) {
	rsp := new(DepartmentGetRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s&id=%d", DepartmentGet, accessToken, departmentid)
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

// 查询部门的所有上级父部门路径
func (c *Client) DepartmentListParent(departmentid int) (*DepartmentListParentRsp, error) {
	rsp := new(DepartmentListParentRsp)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return rsp, err
	}

	url := fmt.Sprintf("%s?access_token=%s&id=%d", DepartmentListParent, accessToken, departmentid)
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
