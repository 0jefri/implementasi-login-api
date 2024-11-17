package users

type UserService interface {
	RegisterNewUser(payload User) error
}

type userService struct {
	userRepo userRepository
}

func NewUserService(usrRepo userRepository) UserService {
	return &userService{userRepo: usrRepo}
}

func (s userService) RegisterNewUser(payload User) error {
	return nil
}
