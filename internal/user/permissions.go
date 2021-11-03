package user

type Permissions struct {
	Ticket TicketPermissions `json:"ticket_permissions"`
	App    AppPermissions    `json:"app"`
}

type BaseResourcePermissions struct {
	View   bool `json:"view"`
	Modify bool `json:"modify"`
}

type TicketPermissions struct {
	Base BaseResourcePermissions

	Comment bool `json:"comment"`
	Close   bool `json:"close"`
	Push    bool `json:"push"`
}

type AppPermissions struct {
	AddUsers          bool
	ChangePermissions bool
}
