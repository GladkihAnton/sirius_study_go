package tests

import (
	"net/http"
	"regexp"
	"strings"
)

type MyHttp struct {
	client http.Client
}

func (m MyHttp) Get(url string) (resp *http.Response, err error) {
	return m.client.Get(url)
}

// MatchContains returns true if the string
// contains the pattern, false otherwise.
type MyHttpI interface {
	Get(url string) (resp *http.Response, err error)
}

func MatchContains(pattern string, src string) bool {
	return strings.Contains(src, pattern)
}

func MatchRegexp(pattern string, src string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(src)
}
