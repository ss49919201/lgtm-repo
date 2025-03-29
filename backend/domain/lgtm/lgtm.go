package lgtm

import (
	"io"
	"net/url"

	"github.com/ss49919201/lgtm-repo/backend/domain/user"
)

type LGTM struct {
	ImageURL url.URL
}

type LGTMImage struct {
	UserID user.UserID
	Image  io.Reader
}
