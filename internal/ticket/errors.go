package ticket

import "fmt"

// ErrInvalidEnumTicketStatus ...
type ErrInvalidEnumTicketStatus string

// Error ...
func (err ErrInvalidEnumTicketStatus) Error() string {
	return fmt.Sprintf("EnumTicketStatus invalid: (%s)", string(err))
}
