package upandgo

import (
	"regexp"
)

type Interceptor struct {
}

type UrlMatcher struct {
	pattern string
	regex   *regexp.Regexp
	params  map[int]string
	method  string
}

type Routers struct {
}
