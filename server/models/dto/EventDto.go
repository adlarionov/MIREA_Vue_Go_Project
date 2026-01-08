package dto

type EventRequestDto struct {
	Event EventDto `json:"event"`
}

type EventDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
