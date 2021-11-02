/*
 * Copyright (c) 2021.
 * Чоч all rights reserved
 */

package user

type Permissions struct {
	Comment bool `json:"comment"`
	Close   bool `json:"close"`
	Push    bool `json:"push"`
	Modify  bool `json:"modify"`
}
