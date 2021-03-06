package user

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// EnumUserType should match 'enum_user_type' enum type.
type EnumUserType uint8

// EnumUserType values.
const (
	// EnumUserTypeAdministrator is the 'Administrator' enum_user_type.
	EnumUserTypeAdministrator EnumUserType = iota + 1
	// EnumUserTypeModerator is the 'Moderator' enum_user_type.
	EnumUserTypeModerator
	// EnumUserTypePrivileged is the 'Privileged' enum_user_type.
	EnumUserTypePrivileged
	// EnumUserTypeRegular is the 'Regular' enum_user_type.
	EnumUserTypeRegular
	// EnumUserTypePending is the 'Pending' enum_user_type.
	EnumUserTypePending
	// EnumUserTypeBlocked is the 'Blocked' enum_user_type.
	EnumUserTypeBlocked
)

// IsAllowed filter for roles.
func (eut EnumUserType) IsAllowed() bool {
	switch eut {
	case EnumUserTypeAdministrator,
		EnumUserTypeModerator,
		EnumUserTypePrivileged,
		EnumUserTypeRegular:
		return true
	}
	return false
}

// String implements fmt.Stringer interface.
func (eut EnumUserType) String() string {
	switch eut {
	case EnumUserTypeAdministrator:
		return "Administrator"
	case EnumUserTypeRegular:
		return "Regular"
	case EnumUserTypeModerator:
		return "Moderator"
	case EnumUserTypePrivileged:
		return "Privileged"
	case EnumUserTypeBlocked:
		return "Blocked"
	case EnumUserTypePending:
		return "Pending"
	}
	return fmt.Sprintf("EnumUserType(%d)", eut)
}

// MarshalText represents EnumUserType as text.
func (eut EnumUserType) MarshalText() ([]byte, error) {
	return []byte(eut.String()), nil
}

// UnmarshalText represents text as EnumUserType.
func (eut *EnumUserType) UnmarshalText(buf []byte) error {
	switch str := string(buf); str {
	case "Administrator":
		*eut = EnumUserTypeAdministrator
	case "Regular":
		*eut = EnumUserTypeRegular
	case "Moderator":
		*eut = EnumUserTypeModerator
	case "Privileged":
		*eut = EnumUserTypePrivileged
	case "Blocked":
		*eut = EnumUserTypeBlocked
	case "Pending":
		*eut = EnumUserTypePending
	default:
		return ErrInvalidEnumUserType(str)
	}
	return nil
}

// Value implements driver.Valuer.
func (eut EnumUserType) Value() (driver.Value, error) {
	return eut.String(), nil
}

// Scan implements sql.Scanner.
func (eut *EnumUserType) Scan(v interface{}) error {
	if buf, ok := v.([]byte); ok {
		return eut.UnmarshalText(buf)
	}
	return ErrInvalidEnumUserType(fmt.Sprintf("%T", v))
}

type User struct {
	ID         uuid.UUID      `json:"id"`
	Login      string         `json:"login"`
	Password   string         `json:"password"` // should be hashed one day :)
	Type       EnumUserType   `json:"type"`
	Admin      bool           `json:"admin"`
	ProfileURL string         `json:"profile_url"`
	Name       sql.NullString `json:"name"`
	AvatarURL  sql.NullString `json:"avatar_url"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  sql.NullTime   `json:"deleted_at"`
}

func (u User) String() string {
	j, _ := json.MarshalIndent(u, "", "    ")
	return string(j)
}

func (u User) GetPermissions() (p Permissions) {
	return
}
