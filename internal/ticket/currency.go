package ticket

import (
	"fmt"
	"strconv"
)

// EnumCurrencyType int value or different currencies
type EnumCurrencyType uint16

const (
	BYN EnumCurrencyType = 54
	USD EnumCurrencyType = 251
	EUR EnumCurrencyType = 93 // not yet allowed by isValid()
)

// IsValid check if currency is allowed in a system.
func (c EnumCurrencyType) IsValid() bool {
	switch c {
	case BYN, USD:
		return true
	}
	return false
}

func (c *EnumCurrencyType) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	if c == nil {
		return fmt.Errorf("nil receiver passed to UnmarshalJSON")
	}

	if ci, err := strconv.ParseUint(string(b), 10, 16); err == nil {
		rtn := EnumCurrencyType(ci)
		if !rtn.IsValid() {
			return fmt.Errorf("invalid currency: %q", ci)
		}

		*c = rtn
		return nil
	}

	if jc, ok := StringToCurrency[string(b)]; ok {
		*c = jc
		return nil
	}
	return fmt.Errorf("invalid currency: %q", string(b))
}

var StringToCurrency = map[string]EnumCurrencyType{
	`"CANDY_WRAPPERS"`: BYN,
	`"MONEY"`:          USD,
}

func (c EnumCurrencyType) String() string {

	switch c {
	case BYN:
		return "CANDY_WRAPPERS"
	case USD:
		return "MONEY"
	case EUR:
		return "NOT_YET_SUPPORTED_CURRENCY"
	}
	return "ERROR_CASE"
}
