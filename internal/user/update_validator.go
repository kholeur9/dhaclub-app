package user

import "github.com/kholeur9/dhaclub-app/internal/apperrors"

func ValidateSingleFieldUpdate(field UpdateDto) error {
	count := 0
	if field.Email != nil {
		count++
	}
	if field.Username != nil {
		count++
	}
	if count == 0 {
		return apperrors.ErrNoFieldsToUpdate
	}
	if count > 1 {
		return apperrors.ErrManyFieldsToUpdate
	}
	return nil
}