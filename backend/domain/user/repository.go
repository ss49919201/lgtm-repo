package user

type Repository interface {
	Save(User) error
	FindByID(UserID) (*User, error)
}
