package ticket

import "fmt"

// ErrInvalidTicketPerk ...
type ErrInvalidTicketPerk string

// Error ...
func (err ErrInvalidTicketPerk) Error() string {
	return fmt.Sprintf("Ticket Perk invalid: (%s)", string(err))
}
