package ticket

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type EnumTicketStatus string

const (
	Draft  EnumTicketStatus = "Draft"
	Active EnumTicketStatus = "Active"
	Closed EnumTicketStatus = "Closed"
)

type Ticket struct {
	ID          uuid.NullUUID    `json:"id"`
	Name        string           `json:"name,omitempty"`
	FullName    string           `json:"full_name,omitempty"`
	Description string           `json:"description,omitempty"`
	Status      EnumTicketStatus `json:"status,omitempty"`
	OwnerID     uint             `json:"owner_id,omitempty"`
	Amount      uint             `json:"amount,omitempty"`
	Price       float64          `json:"price,omitempty"`
	Currency    EnumCurrencyType `json:"currency,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   sql.NullTime     `json:"updated_at"`
	DeletedAt   sql.NullTime     `json:"deleted_at"`
	PublishedAt sql.NullTime     `json:"published_at"`
}

func (t Ticket) IsValid() bool {
	switch t.Status {
	case Draft, Active, Closed:
		return true
	}
	return false
}

func (t Ticket) String() string {
	j, _ := json.MarshalIndent(t, "", "    ")
	return string(j)
}
