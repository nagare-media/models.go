package base

import "net/url"

type URL string

func (u URL) URL() (*url.URL, error) {
	return url.Parse(string(u))
}
