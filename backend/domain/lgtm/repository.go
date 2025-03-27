package lgtm

type Repository interface {
	Save(LGTM) error
	FindByID(LGTMID) (*LGTM, error)
}
