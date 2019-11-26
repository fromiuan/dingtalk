package dingtalk

import (
	"testing"
)

func TestClient(t *testing.T) {
	app_key := ""
	app_secret := ""

	c := NewClient(app_key, app_secret)
	c.SetDebug(true)
	accessToken, err := c.GetAccessToken()
	if err != nil {
		t.Log(err)
	}
	t.Log(accessToken)
}
