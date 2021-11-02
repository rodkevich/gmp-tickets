package ticket

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// EnumTicketStatus should match 'enum_ticket_status' enum type.
type EnumTicketStatus uint8

// IsValid ...
func (ets EnumTicketStatus) IsValid() bool {
	switch ets {
	case EnumTicketStatusDraft,
		EnumTicketStatusActive,
		EnumTicketStatusClosed:
		return true
	}
	return false
}

// EnumTicketStatus values.
const (
	// EnumTicketStatusDraft is the 'Draft' enum_ticket_status.
	EnumTicketStatusDraft EnumTicketStatus = iota + 1
	// EnumTicketStatusActive is the 'Active' enum_ticket_status.
	EnumTicketStatusActive
	// EnumTicketStatusClosed is the 'Closed' enum_ticket_status.
	EnumTicketStatusClosed
)

// String
func (ets EnumTicketStatus) String() string {
	switch ets {
	case EnumTicketStatusDraft:
		return "Draft"
	case EnumTicketStatusActive:
		return "Active"
	case EnumTicketStatusClosed:
		return "Closed"
	}
	return fmt.Sprintf("EnumTicketStatus(%d)", ets)
}

// MarshalText represents EnumTicketStatus as text.
func (ets EnumTicketStatus) MarshalText() ([]byte, error) {
	return []byte(ets.String()), nil
}

// Value implements driver.Valuer interface.
func (ets EnumTicketStatus) Value() (driver.Value, error) {
	return ets.String(), nil
}

// Scan implements sql.Scanner interface.
func (ets *EnumTicketStatus) Scan(v interface{}) error {
	if buf, ok := v.([]byte); ok {
		return ets.UnmarshalText(buf)
	}
	return ErrInvalidEnumTicketStatus(fmt.Sprintf("%T", v))
}

// UnmarshalText represents text as EnumTicketStatus.
func (ets *EnumTicketStatus) UnmarshalText(buf []byte) error {
	switch str := string(buf); str {
	case "Draft":
		*ets = EnumTicketStatusDraft
	case "Active":
		*ets = EnumTicketStatusActive
	case "Closed":
		*ets = EnumTicketStatusClosed
	default:
		return ErrInvalidEnumTicketStatus(str)
	}
	return nil
}

type Ticket struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	FullName    sql.NullString   `json:"full_name"`
	Description string           `json:"description"`
	Status      EnumTicketStatus `json:"status"`
	OwnerID     uuid.UUID        `json:"owner_id"`
	Amount      uint             `json:"amount"`
	Price       sql.NullFloat64  `json:"price"`
	Currency    sql.NullFloat64  `json:"currency"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   sql.NullTime     `json:"deleted_at"`
	PublishedAt sql.NullTime     `json:"published_at"`
}

func (t Ticket) String() string {
	j, _ := json.MarshalIndent(t, "", "    ")
	return string(j)
}

// ErrInvalidEnumTicketStatus ...
type ErrInvalidEnumTicketStatus string

// Error ...
func (err ErrInvalidEnumTicketStatus) Error() string {
	return fmt.Sprintf("EnumTicketStatus invalid: (%s)", string(err))
}
