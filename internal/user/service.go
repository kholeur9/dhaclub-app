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
	email := strings.TrimSpace(u.Email)
	username := strings.TrimSpace(u.Username)
	password := strings.TrimSpace(u.Password)
	if email == "" {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.EmailRequired,
		}
	}
	if username == "" {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.UsernameRequired,
		}
	} else if len(username) <= 1 {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.UsernameShort,
		}
	}
	if password == "" {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.PasswordRequired,
		}
	} else if len(password) < 8 {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.PasswordShort,
		}
	}
	userExists, err := us.userStore.FindConflicts(&email, &username)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	if userExists != nil {
		if userExists.Email == email {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: apperrors.EmailAlreadyExists,
			}
		}
		if userExists.Username == username {
			return nil, &apperrors.ServiceError{
				Type:    apperrors.CONFLICT,
				Message: apperrors.UsernameAlreadyExists,
			}
		}
	}
	passwordhashed, err := us.secure.PasswordHash(password)
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	newUser := User{
		ID:       uuid.New().String(),
		Email:    email,
		Username: username,
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

func (us *UserService) Authenticate(identifier, password string) (*User, error) {
	if identifier == "" {
		return nil, apperrors.ErrIdentifierEmpty
	}
	if password == "" {
		return nil, apperrors.ErrPasswordEmpty
	}
	var user *User
	var err error
	if strings.Contains(identifier, "@") {
		user, err = us.userStore.GetUserByEmail(identifier)
	} else {
		user, err = us.userStore.GetUserByUsername(identifier)
	}
	if err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, apperrors.ErrAccountNotActive
	}
	passwordMatch := us.secure.PasswordCompare(user.Password, password)
	if !passwordMatch {
		return nil, apperrors.ErrPasswordWrong
	}
	return user, err
}

func (us *UserService) UserUpdate(id string, ud UpdateDto) (*UpdateResponseDto, error) {
	err := ValidateSingleFieldUpdate(ud)
	if errors.Is(err, apperrors.ErrManyFieldsToUpdate) {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: apperrors.OneFieldToUpdate,
		}
	} else if errors.Is(err, apperrors.ErrNoFieldsToUpdate) {
		return nil, &apperrors.ServiceError{
			Type: apperrors.VALIDATION,
			Message: apperrors.NoDataToUpdate,
		}
	}
	getUser, err := us.userStore.GetUserById(id)
	if errors.Is(err, apperrors.ErrUserNotFound) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.NOT_FOUND,
			Message: apperrors.UserNotFoundMessage,
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
	if ud.Email != nil {
		email := strings.TrimSpace(*ud.Email)
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
	if ud.Username != nil {
		username := strings.TrimSpace(*ud.Username)
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

func (us *UserService) UserUpdatePassword(id string, password UpdatePasswordDto) (*UpdatePasswordResponse, error) {
	return &UpdatePasswordResponse{}, nil
}