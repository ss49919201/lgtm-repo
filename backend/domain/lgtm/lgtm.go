package lgtm

import (
	"io"
	"net/url"

	"github.com/ss49919201/lgtm-repo/backend/domain/errs"
)

type LGTMID struct {
	value int
}

func NewLGTMID(v int) (LGTMID, error) {
	if !(v > 0) {
		return LGTMID{}, errs.NewDomainError("lgtm must be greater than 0")
	}

	return LGTMID{v}, nil
}

type LGTM struct {
	ID       LGTMID
	ImageURL url.URL
}

type LGTMImage struct {
	Image io.Reader
}
