package dto

type BookingRequestDto struct {
	BookingDto
}

type BookingDto struct {
	Price  int    `json:"price"`
	Status string `json:"status"`
}
