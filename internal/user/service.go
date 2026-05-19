package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kholeur9/dhaclub-app/internal/apperrors"
	"github.com/kholeur9/dhaclub-app/internal/helpers"
)

type UserService struct {
	userStore UserStore
	secure helpers.PasswordSecure
}

func NewUserService(userStore UserStore, secure helpers.PasswordSecure) *UserService{
	return &UserService{
		userStore: userStore,
		secure: secure,
	}
}

func (us *UserService) CreateUser(u CreateUserDto) (*CreateUserResponseDto, error) {
	fmt.Println("Service create user :", u)
	if u.Email == "" {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: "Email is required",
		}
	}
	if u.Username == "" {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: "Username is required",
		}
	} else if len(u.Username) <= 1 {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: "Username too short, enter more one caractere.",
		}
	} 
	if u.Password == "" {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: "Password is required",
		}
	} else if len(u.Password) < 8 {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: "Password would have 8 caracteres minimum.",
		}
	}
	passwordhashed, err := us.secure.PasswordHash(u.Password)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type: apperrors.INTERNAL,
			Message: "Internal server error",
		}
	}
	newUser := User{
		ID: uuid.New().String(),
		Email: u.Email,
		Username: u.Username,
		Password: passwordhashed,
	}
	_, err = us.userStore.Create(newUser)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type: apperrors.INTERNAL,
			Message: "",
		}
	}
	return nil, nil
}
