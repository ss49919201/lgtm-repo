package errs

type DomainError struct {
	msg string
}

func (d DomainError) Error() string {
	return d.msg
}

func NewDomainError(msg string) error {
	return DomainError{
		msg: msg,
	}
}
