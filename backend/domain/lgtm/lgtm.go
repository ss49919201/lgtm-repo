package lgtm

import (
	"io"
	"net/url"
)

type LGTMID struct {
	value int
}

type LGTM struct {
	ID       LGTMID
	ImageURL url.URL
}

type LGTMImage struct {
	Image io.Reader
}
