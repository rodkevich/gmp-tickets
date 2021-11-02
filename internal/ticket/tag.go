package ticket

type Tag struct {
	Name       string `json:"name"`
	Searchable bool   `json:"searchable"`
}

type Tags struct {
	Count int   `json:"count"`
	List  []Tag `json:"list"`
}
