package user

import "github.com/ss49919201/lgtm-repo/backend/domain/errs"

type UserID struct {
	value int
}

func NewUserID(v int) (UserID, error) {
	if !(v > 0) {
		return UserID{}, errs.NewDomainError("user id must be greater than 0")
	}

	return UserID{v}, nil
}

type User struct {
	ID UserID
}
