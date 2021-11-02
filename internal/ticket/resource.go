package ticket

import (
	"github.com/rodkevich/gmp-tickets/internal/user"
)

type RtnFull struct {
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
