![](https://img.alicdn.com/tfs/TB1bB9QKpzqK1RjSZFoXXbfcXXa-576-96.png)



#dingtalk

企业内部应用SDK，内置多重缓存方式本地存储、memcache、redis，关于企业内部应用接口的具体说明可以参考[官方文档](https://ding-doc.dingtalk.com/doc#/serverapi2/gh60vz)

## 企业内部应用

- 身份验证(auth)
- 通讯录管理
 - 用户管理(user)
 - 部门管理(department) 

	
## Install

```bash
go get github.com/fromiuan/dingtalk
```

## Usage

```bash
	
	package main

	import (
		"fmt"
	
		"github.com/fromiuan/dingtalk"
	)
	
	var (
		appkey    = ""
		appsecret = ""
	)
	
	func main() {
		// 扫描登陆第三方、钉钉内部登录第三方、密码登录第三方不需要appkey，appsecret
		client := dingtalk.NewClient(appkey, appsecret)
		client.SetDebug(true)
	
		accessToken, err := client.GetAccessToken()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(accessToken)
	
		// appId、appSecret 为扫描登陆的参数
		userAuth, err := client.GetUserInfoByCode("code", "appId", "appSecret")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(userAuth)
	
		adminList, err := client.UserGetAdmin()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(adminList)
	}


```