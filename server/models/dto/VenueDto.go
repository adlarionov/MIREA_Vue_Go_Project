package dto

type VenueRequestDto struct {
	Venue VenueDto
}

type VenueDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity"`
	Address     string `json:"address"`
}
