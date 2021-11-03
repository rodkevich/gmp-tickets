package ticket

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rodkevich/gmp-tickets/internal/user"
	"io"
)

type CreationRequest struct {
	ID          uuid.UUID        `json:"-"`
	Name        string           `json:"name"`
	FullName    *string          `json:"full_name"`
	Description string           `json:"description"`
	Status      EnumTicketStatus `json:"status"`
	OwnerID     uuid.UUID        `json:"owner_id"`
	Amount      uint             `json:"amount"`
	Price       *float64         `json:"price"`
	Currency    uint             `json:"currency"`
}

func Bind(body io.ReadCloser, b *CreationRequest) error {
	return json.NewDecoder(body).Decode(b)
}

func Parse(req *CreationRequest) *Ticket {
	t := Ticket{
		ID:   req.ID,
		Name: req.Name,
		FullName: sql.NullString{
			String: *req.FullName,
			Valid:  true,
		},
		Description: req.Description,
		Status:      req.Status,
		OwnerID:     req.OwnerID,
		Amount:      req.Amount,
		Price: sql.NullFloat64{
			Float64: *req.Price,
			Valid:   true,
		},
		Currency: sql.NullInt64{
			Int64: int64(req.Currency),
			Valid: true,
		},
	}
	return &t
}

type Rtn struct {
	Base  Ticket    `json:"base"`
	Price Price     `json:"price"`
	Owner user.User `json:"owner"`
	Extra *Extra    `json:"extra,omitempty"`
}

type Extra struct {
	Description *string `json:"description,omitempty"`
	Photos      *Photos `json:"photos,omitempty"`
}

type Price struct {
	Amount   int              `json:"amount"`
	Currency EnumCurrencyType `json:"currency"`
}

type Description struct {
	Length int    `json:"length"`
	Text   string `json:"text"`
}
