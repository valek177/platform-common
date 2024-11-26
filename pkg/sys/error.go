package sys

import (
	"github.com/pkg/errors"

	"github.com/valek177/platform-common/pkg/sys/codes"
)

type commonError struct {
	msg  string
	code codes.Code
}

// NewCommonError returns common error
func NewCommonError(msg string, code codes.Code) *commonError {
	return &commonError{msg, code}
}

func (r *commonError) Error() string {
	return r.msg
}

func (r *commonError) Code() codes.Code {
	return r.code
}

// IsCommonError checks error is common
func IsCommonError(err error) bool {
	var ce *commonError
	return errors.As(err, &ce)
}

// GetCommonError returns common error
func GetCommonError(err error) *commonError {
	var ce *commonError
	if !errors.As(err, &ce) {
		return nil
	}

	return ce
}
