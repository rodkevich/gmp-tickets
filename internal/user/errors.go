package user

import "fmt"

// ErrInvalidEnumUserType is the invalid EnumUserType error.
type ErrInvalidEnumUserType string

// Error implements the builtin.error.
func (err ErrInvalidEnumUserType) Error() string {
	return fmt.Sprintf("invalid EnumUserType(%s)", string(err))
}
