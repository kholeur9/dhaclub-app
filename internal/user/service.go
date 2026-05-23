package user

import (
	//"fmt"

	//"errors"

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
	var newUser User
	if u.Email == "" {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: apperrors.EmailRequired,
		}
	}
	if u.Username == "" {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: apperrors.UsernameRequired,
		}
	} else if len(u.Username) <= 1 {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: apperrors.UsernameShort,
		}
	} 
	if u.Password == "" {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: apperrors.PasswordRequired,
		}
	} else if len(u.Password) < 8 {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: apperrors.PasswordShort,
		}
	}
	userExists, err := us.userStore.UserExistsByEmail(u.Email)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type: apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	if userExists {
		return nil, &apperrors.ServiceError{
			Type: apperrors.CONFLICT,
			Message: apperrors.UserAlreadyExists,
		}
	} else {
		passwordhashed, err := us.secure.PasswordHash(u.Password)
		if err != nil {
			return nil, &apperrors.ServiceError{
				Type: apperrors.INTERNAL,
				Message: apperrors.ErrInternalServerErrorMessage,
			}
		}
		newUser = User{
			ID: uuid.New().String(),
			Email: u.Email,
			Username: u.Username,
			Password: passwordhashed,
		}

	}
	user, err := us.userStore.Create(newUser)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type: apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	return &CreateUserResponseDto{
		Message: "Successfuly",
		Data: CreateUserResponse{
			ID: user.ID,
			Email: user.Email,
			Username: user.Username,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}
