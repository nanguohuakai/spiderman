package dto

type OwnedProject struct {
	ID       string `json:"id"`
	Level    string `json:"level"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	URL      string `json:"url"`
	Readable uint8  `json:"readable"`
}
