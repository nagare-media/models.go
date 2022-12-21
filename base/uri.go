package base

import "net/url"

type URI string

func (u URI) URL() (*url.URL, error) {
	return url.Parse(string(u))
}
