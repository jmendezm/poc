package sites

import (
	"technical_test/domain"
)

func HandleError(err error) *domain.CustomError {
	if ce, ok := err.(*domain.CustomError); ok {
		return ce
	}
	return &domain.ErrInternalServerError
}
