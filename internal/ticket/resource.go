package ticket

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"time"
)

type CreationRequest struct {
	ID          uuid.UUID      `json:"-"`
	OwnerID     uuid.UUID      `json:"owner_id"`
	NameShort   string         `json:"name_short"`
	NameExt     string         `json:"name_ext"`
	Description string         `json:"description"`
	Amount      int16          `json:"amount"`
	Price       float64        `json:"price"`
	Currency    int16          `json:"currency"`
	Active      bool           `json:"active"`
	Perk        EnumTicketPerk `json:"perk"`
	PublishedAt *time.Time     `json:"published_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at"`
}

func Bind(body io.ReadCloser, b *CreationRequest) error {
	return json.NewDecoder(body).Decode(b)
}

func Parse(req *CreationRequest) *Ticket {
	t := Ticket{
		ID:          req.ID,
		OwnerID:     req.OwnerID,
		NameShort:   req.NameShort,
		NameExt:     req.NameExt,
		Description: req.Description,
		Amount:      req.Amount,
		Price:       req.Price,
		Currency:    req.Currency,
		Active:      req.Active,
		Perk:        req.Perk,
		PublishedAt: req.PublishedAt,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
		DeletedAt:   req.DeletedAt,
	}
	return &t
}

type Subject struct {
	ID          uuid.UUID  `json:"id"`
	OwnerID     uuid.UUID  `json:"owner_id"`
	NameShort   string     `json:"name_short"`
	Active      bool       `json:"active"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type Rtn struct {
	Subject Subject `json:"base"`
	Price   Price   `json:"price"`
	Extra   *Extra  `json:"extra,omitempty"`
}

func RtnPrepare(t Ticket, ext bool, p *Photos) (response Rtn) {
	response = Rtn{
		Subject: Subject{
			ID:          t.ID,
			OwnerID:     t.OwnerID,
			NameShort:   t.NameShort,
			Active:      t.Active,
			PublishedAt: t.PublishedAt,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
			DeletedAt:   t.DeletedAt,
		},
		Price: Price{
			Amount:   t.Amount,
			Price:    t.Price,
			Currency: t.Currency,
		},
	}
	switch ext {
	case true:
		response.Extra = &Extra{
			Description: &t.Description,
			Photos:      p,
			Perk:        t.Perk,
		}
	}

	return
}

type Extra struct {
	Description *string        `json:"description,omitempty"`
	Photos      *Photos        `json:"photos,omitempty"`
	Perk        EnumTicketPerk `json:"perk"`
}

type Price struct {
	Amount   int16   `json:"amount"`
	Price    float64 `json:"price"`
	Currency int16   `json:"currency"`
}

type Description struct {
	Length int    `json:"length"`
	Text   string `json:"text"`
}
