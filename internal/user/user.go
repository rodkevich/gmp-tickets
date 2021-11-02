/*
 * Copyright (c) 2021.
 * Чоч all rights reserved
 */

package user

import (
	"encoding/json"
	"time"
)

type EnumUserType string

// NOTE: should match db - enum_user_type
const (
	Administrator EnumUserType = "Administrator"
	Regular       EnumUserType = "Regular"
	Moderator     EnumUserType = "Moderator"
	Privileged    EnumUserType = "Privileged"
	Blocked       EnumUserType = "Blocked"
	Pending       EnumUserType = "Pending"
)

type User struct {
	ID         uint         `json:"id,omitempty"`
	Login      string       `json:"login,omitempty"`
	AvatarUrl  string       `json:"avatar_url,omitempty"`
	Url        string       `json:"url,omitempty"`
	Name       string       `json:"name,omitempty"`
	Type       EnumUserType `json:"type,omitempty"`
	Admin      bool         `json:"admin,omitempty"`
	CreatedAt  time.Time    `json:"created_at"`
	ModifiedAt *time.Time   `json:"modified_at,omitempty"`
}

func (u User) IsValid() bool {
	switch u.Type {
	case Administrator, Regular, Moderator, Privileged:
		return true
	}
	return false
}

func (u User) String() string {
	j, _ := json.MarshalIndent(u, "", "    ")
	return string(j)
}

func (u User) GetPermissions() (p Permissions) {
	return
}
