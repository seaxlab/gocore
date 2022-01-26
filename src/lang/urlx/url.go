package urlx

import "net/url"

// spy 2022/1/26

// GetDomain
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name
// will get www.aspxfans.com
func GetDomain(rawUrl string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	return stUrl.Hostname()
}

// GetPort
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name
// will get 8080
func GetPort(rawUrl string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	return stUrl.Port()
}
