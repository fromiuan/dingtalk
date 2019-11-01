package dingtalk

type Client struct {
	AppId     string
	AppSecret string
	Debug     bool
}

func NewClient(appid, appsecret string) *Client {
	return &Client{
		AppId:     appid,
		AppSecret: appsecret,
	}
}

func (c *Client) SetDebug(b bool) {
	c.Debug = b
}
