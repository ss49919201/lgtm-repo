package lgtm

import (
	"io"
	"net/url"
)

type LGTM struct {
	ImageURL url.URL
}

type LGTMImage struct {
	Image io.Reader
}
