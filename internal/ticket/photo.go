package ticket

import (
	"database/sql"

	"github.com/google/uuid"
)

type Photo struct {
	ID          uuid.NullUUID
	TicketID    uuid.UUID
	IsMainPhoto sql.NullBool
	Presented   sql.NullBool
	MimeType    sql.NullString
	SizeKb      uint
	LinkAddress string
}

func (p Photo) String() string {
	return p.LinkAddress
}

type Photos struct {
	Main       Photo            `json:"main_photo"`
	Additional PhotosAdditional `json:"additional_photos"`
}

type PhotosAdditional struct {
	Count int           `json:"count"`
	Link  string        `json:"link"`
	Batch map[int]Photo `json:"batch"`
}

func (p PhotosAdditional) String() string {
	return p.Link
}
