package user

import (
	"fmt"
	"strings"
	"time"

	"errors"

	"github.com/google/uuid"
	"github.com/kholeur9/dhaclub-app/internal/apperrors"
	"github.com/kholeur9/dhaclub-app/internal/helpers"
)

type UserService struct {
	userStore UserStore
	secure    helpers.PasswordSecure
}

func NewUserService(userStore UserStore, secure helpers.PasswordSecure) *UserService {
	return &UserService{
		userStore: userStore,
		secure:    secure,
	}
}

func (us *UserService) CreateUser(u CreateUserDto) (*CreateUserResponseDto, error) {
	//var newUser User
	if u.Email == "" {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.EmailRequired,
		}
	}
	if u.Username == "" {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.UsernameRequired,
		}
	} else if len(u.Username) <= 1 {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.UsernameShort,
		}
	}
	if u.Password == "" {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.PasswordRequired,
		}
	} else if len(u.Password) < 8 {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.PasswordShort,
		}
	}
	userExists, err := us.userStore.FindConflicts(&u.Email, &u.Username)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	if userExists != nil {
		if userExists.Email == u.Email {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: apperrors.EmailAlreadyExists,
			}
		}
		if userExists.Username == u.Username {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: apperrors.UsernameAlreadyExists,
			}
		}
	}
	passwordhashed, err := us.secure.PasswordHash(u.Password)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	newUser := User{
		ID:       uuid.New().String(),
		Email:    u.Email,
		Username: u.Username,
		Password: passwordhashed,
	}
	user, err := us.userStore.Create(newUser)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	return &CreateUserResponseDto{
		Message: "Successfuly",
		Data: CreateUserResponse{
			ID:        user.ID,
			Email:     user.Email,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}

func (us *UserService) UserUpdate(id string, eu UpdateDto) (*UpdateResponseDto, error) {
	if eu.Email == nil && eu.Username == nil {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: "No data to update",
		}
	}
	getUser, err := us.userStore.GetUserById(id)
	if errors.Is(err, apperrors.ErrUserNotFound) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: apperrors.ErrUserNotFoundMessage,
		}
	}
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	var sendNewData = UpdateResponse{
		ID: id,
		UpdatedAt: time.Now(),
	}
	err = ValidateSingleFieldUpdate(eu)
	if errors.Is(err, apperrors.ErrManyFieldsToUpdate) {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: apperrors.OneFieldToUpdate,
		}
	}
	if eu.Email != nil {
		email := strings.TrimSpace(*eu.Email)
		if email == "" {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.VALIDATION,
				Message: apperrors.EmailRequired,
			}
		}
		if getUser.Email == email {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: apperrors.EmailNotChanged,
			}
		}
		emailUsed, err := us.userStore.FindConflicts(&email, nil)
		if err != nil {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.INTERNAL,
				Message: apperrors.ErrInternalServerErrorMessage,
			}
		}
		if emailUsed != nil {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: apperrors.EmailAlreadyExists,
			}
		}
		sendNewData.Email = &email
	}
	if eu.Username != nil {
		username := strings.TrimSpace(*eu.Username)
		if username == "" {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.VALIDATION,
				Message: apperrors.UsernameRequired,
			}
		}
		if len(username) <= 1 {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.VALIDATION,
				Message: apperrors.UsernameShort,
			}
		}
		if getUser.Username == username {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: apperrors.UsernameNotChanged,
			}
		}
		usernameUsed, err := us.userStore.FindConflicts(nil, &username)
		if err != nil {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.INTERNAL,
				Message: apperrors.ErrInternalServerErrorMessage,
			}
		}
		if usernameUsed != nil {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: apperrors.UsernameAlreadyExists,
			}
		}
		sendNewData.Username = &username
	}
	userUpdate, err := us.userStore.Update(sendNewData)
	fmt.Println("update err:", err)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	return &UpdateResponseDto{
		Message: "Successfuly",
		Data: UpdateResponse{
			ID:        userUpdate.ID,
			Email:     &userUpdate.Email,
			Username:  &userUpdate.Username,
			UpdatedAt: *userUpdate.UpdatedAt,
		},
	}, nil
}
