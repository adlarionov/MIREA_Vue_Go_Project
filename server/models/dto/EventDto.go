package dto

type EventRequestDto struct {
	OwnerEmail string   `json:"owner"`
	Event      EventDto `json:"event"`
}

type EventDto struct {
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Address      string  `json:"address"`
	City         string  `json:"city"`
	Capacity     int     `json:"capacity"`
	PricePerHour float32 `json:"price_per_hour"`
}
