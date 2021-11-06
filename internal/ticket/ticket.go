package ticket

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// EnumTicketPerk should match 'enum_ticket_status' enum type.
type EnumTicketPerk uint16

// IsValid ...
func (ets EnumTicketPerk) IsValid() bool {
	switch ets {
	case EnumTicketPerkDraft,
		EnumTicketPerkRegular,
		EnumTicketPerkPremium:
		return true
	}
	return false
}

// EnumTicketPerk values.
const (
	// EnumTicketPerkDraft is the 'Draft' enum_tickets_perk_type.
	EnumTicketPerkDraft EnumTicketPerk = iota + 1
	// EnumTicketPerkRegular is the 'Regular' enum_tickets_perk_type.
	EnumTicketPerkRegular
	// EnumTicketPerkPremium is the 'Premium' enum_tickets_perk_type.
	EnumTicketPerkPremium
	// EnumTicketPerkPromoted is the 'Promoted' enum_tickets_perk_type.
	EnumTicketPerkPromoted
)

// String
func (ets EnumTicketPerk) String() string {
	switch ets {
	case EnumTicketPerkDraft:
		return "Draft"
	case EnumTicketPerkRegular:
		return "Regular"
	case EnumTicketPerkPremium:
		return "Premium"
	case EnumTicketPerkPromoted:
		return "Promoted"
	}
	return fmt.Sprintf("Ticket Perk is invalid(%d)", ets)
}

// MarshalText represents EnumTicketPerk as text.
func (ets EnumTicketPerk) MarshalText() ([]byte, error) {
	return []byte(ets.String()), nil
}

// Value implements driver.Valuer interface.
func (ets EnumTicketPerk) Value() (driver.Value, error) {
	return ets.String(), nil
}

// Scan implements sql.Scanner interface.
func (ets *EnumTicketPerk) Scan(v interface{}) error {
	if buf, ok := v.([]byte); ok {
		return ets.UnmarshalText(buf)
	}
	return ErrInvalidTicketPerk(fmt.Sprintf("%T", v))
}

// UnmarshalText represents text as EnumTicketPerk.
func (ets *EnumTicketPerk) UnmarshalText(buf []byte) error {
	switch str := string(buf); str {
	case "Draft":
		*ets = EnumTicketPerkDraft
	case "Regular":
		*ets = EnumTicketPerkRegular
	case "Premium":
		*ets = EnumTicketPerkPremium
	case "Promoted":
		*ets = EnumTicketPerkPremium
	default:
		return ErrInvalidTicketPerk(str)
	}
	return nil
}

type Ticket struct {
	ID          uuid.UUID      `json:"id"`           // id
	OwnerID     uuid.UUID      `json:"owner_id"`     // owner_id
	NameShort   string         `json:"name_short"`   // name
	NameExt     string         `json:"name_ext"`     // name_ext
	Description string         `json:"description"`  // description
	Amount      int16          `json:"amount"`       // amount
	Price       float64        `json:"price"`        // price
	Currency    int16          `json:"currency"`     // currency
	Active      bool           `json:"active"`       // active
	Perk        EnumTicketPerk `json:"perk"`         // perk
	PublishedAt *time.Time     `json:"published_at"` // published_at
	CreatedAt   time.Time      `json:"created_at"`   // created_at
	UpdatedAt   time.Time      `json:"updated_at"`   // updated_at
	DeletedAt   *time.Time     `json:"deleted_at"`   // deleted_at
}

func (t Ticket) String() string {
	j, _ := json.MarshalIndent(t, "", "    ")
	return string(j)
}
