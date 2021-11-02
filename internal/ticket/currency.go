package ticket

// EnumCurrencyType int value or different currencies
type EnumCurrencyType uint

const (
	BYN EnumCurrencyType = 54
	USD EnumCurrencyType = 251
	EUR EnumCurrencyType = 93 // not yet allowed by isValid()
)

// IsValid can check if currency is allowed in a system.
func (c EnumCurrencyType) IsValid() bool {

	switch c {
	case BYN, USD:
		return true
	}
	return false
}

// // CurrencyToString ...
// func CurrencyToString(c Currency) string {
//
// 	switch c {
// 	case BYN:
// 		return "Candy wrappers"
// 	case USD:
// 		return "Money"
// 	case EUR:
// 		return "Not yet supported"
// 	}
// 	return "ErrorCase"
// }
