package urlx

import (
	"errors"
	"net/url"
)

// spy 2022/1/26

func BuildParam(params map[string]string) string {
	if params == nil {
		return ""
	}
	query := url.Values{}
	for k, v := range params {
		query.Add(k, v)
	}
	return query.Encode()
}

// GetParam
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key="page" will get "1"
func GetParam(rawUrl, key string) (string, error) {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	m := stUrl.Query()
	if v, ok := m[key]; ok && len(v) > 0 {
		return v[0], nil
	}
	return "", errors.New("no param")
}

// GetParams
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page will get [1 2]
func GetParams(rawUrl, key string) ([]string, error) {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	m := stUrl.Query()
	if v, ok := m[key]; ok {
		return v, nil
	}
	return nil, errors.New("no param")
}

// GetParamsAsMap
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name will get map[boardID:[520] page:[1 2]]
func GetParamsAsMap(rawUrl string) (map[string][]string, error) {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	m := stUrl.Query()
	return m, nil
}

// AddParam
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page value=3
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2&page=3#name
func AddParam(rawUrl, key, value string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	m.Add(key, value)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// AddParams add params
func AddParams(rawUrl string, params map[string]string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	for k, v := range params {
		m.Add(k, v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// DelParam
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520#name
func DelParam(rawUrl, key string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	m.Del(key)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// DelParams delete multi params
func DelParams(rawUrl string, keys []string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	for _, v := range keys {
		m.Del(v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// SetParam
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page value=3
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=3#name
func SetParam(rawUrl, key, value string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	m.Set(key, value)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// SetParams set multi params
func SetParams(rawUrl string, params map[string]string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	for k, v := range params {
		m.Set(k, v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// QueryGetParam get the specified key parameter from query string
// rawQuery is encoded query values without '?'. key is parameter name
// e.g. if query is "a=dog&b=tiger" and key is "a" will return dog
func QueryGetParam(rawQuery, key string) (string, error) {
	queries, err := url.ParseQuery(rawQuery)
	if err != nil {
		return "", err
	}

	if v, ok := queries[key]; ok && len(v) > 0 {
		return v[0], nil
	}
	return "", errors.New("no param")
}

// QueryGetParams get the specified key parameters from query string
// rawQuery is encoded query values without '?'. key is parameter name
// e.g. if query is "a=dog&a=cat&b=tiger" and key is "a" will return [dog cat]
func QueryGetParams(rawQuery, key string) ([]string, error) {
	queries, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, err
	}

	if v, ok := queries[key]; ok && len(v) > 0 {
		return v, nil
	}
	return nil, errors.New("no param")
}
