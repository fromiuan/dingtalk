package dingtalk

const ROOTURL = "https://oapi.dingtalk.com"

var (
	// 身份验证
	GetUserInfoBycode = ROOTURL + "/sns/getuserinfo_bycode"
	GetUserInfo       = ROOTURL + "/user/getuserinfo"
	GetToken          = ROOTURL + "/gettoken"
	SSOGetToken       = ROOTURL + "/sso/gettoken"
	SSOGetUserInfo    = ROOTURL + "/sso/getuserinfo"

	// 通讯录管理-用户管理
	UserCreate             = ROOTURL + "/user/create"
	UserUpdate             = ROOTURL + "/user/update"
	UserDelete             = ROOTURL + "/user/delete"
	UserGet                = ROOTURL + "/user/get"
	UserGetDeptMember      = ROOTURL + "/user/getDeptMember"
	UserSimpleList         = ROOTURL + "/user/simplelist"
	UserListByPage         = ROOTURL + "/user/listbypage"
	UserGetAdmin           = ROOTURL + "/user/get_admin"
	UserGetAdminScope      = ROOTURL + "/topapi/user/get_admin_scope"
	UserGetUseridByUnionid = ROOTURL + "/user/getUseridByUnionid"
	UserGetByMobile        = ROOTURL + "/user/get_by_mobile"
	UserGetOrgUserCount    = ROOTURL + "/user/get_org_user_count"
	UserGetInactive        = ROOTURL + "/topapi/inactive/user/get"

	// 通讯录管理-部门管理
	DepartmentCreate          = ROOTURL + "/department/create"
	DepartmentUpdate          = ROOTURL + "/department/update"
	DepartmentDelete          = ROOTURL + "/department/delete"
	DepartmentListIds         = ROOTURL + "/department/list_ids"
	DepartmentList            = ROOTURL + "/department/list"
	DepartmentGet             = ROOTURL + "/department/get"
	DepartmentListParent      = ROOTURL + "/department/list_parent_depts_by_dept"
	DepartmentListParentDepts = ROOTURL + "/department/list_parent_depts"

	// 通讯录管理-角色管理

)
