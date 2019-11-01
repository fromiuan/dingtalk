package lib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
)

func UserSignatur(timestamp, appSecret string) string {
	h := hmac.New(sha256.New, []byte(appSecret))
	h.Write([]byte(timestamp))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func UrlEncode(vMap map[string]string) string {
	u := url.Values{}
	for k, v := range vMap {
		u.Set(k, v)
	}
	return u.Encode()
}
