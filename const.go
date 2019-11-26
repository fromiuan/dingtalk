package dingtalk

const ROOTURL = "https://oapi.dingtalk.com"

var (
	// 身份验证
	GetUserInfoBycode = ROOTURL + "/sns/getuserinfo_bycode"
	GetUserInfo       = ROOTURL + "/user/getuserinfo"
	GetToken          = ROOTURL + "/gettoken"
	SSOGetToken       = ROOTURL + "/sso/gettoken"
	SSOGetUserInfo    = ROOTURL + "/sso/getuserinfo"

	// 通讯录管理
	UserCreate        = ROOTURL + "/user/create"
	UserUpdate        = ROOTURL + "/user/update"
	UserDelete        = ROOTURL + "/user/delete"
	UserGet           = ROOTURL + "/user/get"
	UserGetDeptMember = ROOTURL + "/user/getDeptMember"
)
