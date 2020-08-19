package types

type Monster struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Hunted bool `json:"hunted"`
}