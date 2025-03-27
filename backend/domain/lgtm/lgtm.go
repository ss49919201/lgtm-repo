package lgtm

import "net/url"

type LGTMID struct {
	value int
}

type LGTM struct {
	ID       LGTMID
	ImageURL url.URL
}
