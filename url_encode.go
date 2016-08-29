package goutils

import "net/url"

// UrlEncoded encodes a string like Javascript's encodeURIComponent()
func EncodeUrlParam(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
