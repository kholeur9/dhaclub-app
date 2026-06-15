package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kholeur9/dhaclub-app/internal/apperrors"
	"github.com/kholeur9/dhaclub-app/internal/user"
)

type AuthService struct {
	userService *user.UserService
}

func NewAuthService(userService *user.UserService) *AuthService{
	return &AuthService{
		userService: userService,
	}
}

func (as *AuthService) LoginUser(lu user.LoginUserDto) (*user.LoginUserResponse, error) {
	fmt.Println("Login Service: ", lu)
	identifier := strings.TrimSpace(lu.Identifer)
	password := strings.TrimSpace(lu.Password)
	userFound, err := as.userService.Authenticate(identifier, password)
	if errors.Is(err, apperrors.ErrIdentifierEmpty) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.IdentifierRequired,
		}
	}
	if errors.Is(err, apperrors.ErrPasswordEmpty) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.VALIDATION,
			Message: apperrors.PasswordRequired,
		}
	}
	if errors.Is(err, apperrors.ErrUserNotFound) || errors.Is(err, apperrors.ErrPasswordWrong) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.UNAUTHORIZED,
			Message: apperrors.InvalidCredentials,
		}
	}
	if errors.Is(err, apperrors.ErrAccountNotActive) {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.UNAUTHORIZED,
			Message: apperrors.AccountNotActive,
		}
	}
	if err != nil {
		return nil, &apperrors.ServiceError{
			Type:    apperrors.INTERNAL,
			Message: apperrors.ErrInternalServerErrorMessage,
		}
	}
	return &user.LoginUserResponse{
		Message: "Successfuly",
		Data: user.LoginResponse{
			ID:        userFound.ID,
			Email:     userFound.Email,
			Username:  userFound.Username,
			IsActive:  userFound.IsActive,
			CreatedAt: userFound.CreatedAt,
			UpdatedAt: userFound.UpdatedAt,
		},
	}, nil
}