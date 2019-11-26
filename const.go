package dingtalk

const ROOTURL = "https://oapi.dingtalk.com"

var (
	getuserinfo_bycode = ROOTURL + "/sns/getuserinfo_bycode"
	getuserinfo        = ROOTURL + "/user/getuserinfo"
	gettoken           = ROOTURL + "/gettoken"
	sso_gettoken       = ROOTURL + "/sso/gettoken"
)
