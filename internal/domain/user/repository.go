package user

type Repository interface {
	Add(user *User) error
	//FindById(id string) (User, error)
	//FindByEmail(email string) (User, error)
}
