package upandgo

import (
	"net/http"
	"net/textproto"
)

type Headers http.Header

func (h Headers) Add(key, value string) {
	textproto.MIMEHeader(h).Add(key, value)
}
