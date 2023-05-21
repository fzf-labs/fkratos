package errorx

import "github.com/go-kratos/kratos/v2/errors"

func New(code int, reason, message string) *errors.Error {
	return errors.New(code, reason, message)
}
