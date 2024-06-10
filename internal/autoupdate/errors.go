package autoupdate

import (
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

type permissionDeniedError struct {
	err error
}

func (e permissionDeniedError) Error() string {
	return fmt.Sprintf("permissoin denied: %v", e.err)
}

func (e permissionDeniedError) Type() string {
	return "permission_denied"
}

type invalidInputError struct {
	msg string
}

func (e invalidInputError) Error() string {
	return e.msg
}

func (e invalidInputError) Type() string {
	return "invalid_input"
}

type notExistError struct {
	key dskey.Key
}

func (e notExistError) Error() string {
	return fmt.Sprintf("%s does not exist", e.key)
}

func (e notExistError) Type() string {
	return "not_exist"
}
