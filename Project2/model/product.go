package model

type Product struct {
	ID    string `json:"ID",bun:"ID"`
	Name  string `json:"Name",bun:"Name"`
	Price int    `json:"Price",bun:"Price"`
}
