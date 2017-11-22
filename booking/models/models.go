package models

type(
	BookingsResource struct {
		Data    []Booking	`json:"data"`
	}
	BookingResource struct{
		Data     Booking         `json:"data"`
	}
	Booking struct {
		UserId       string	`json:"userid"`
		ShowTimeId   string     `json:"showtimeid"`
		Movies   []string       `json:"movies"`
	}
)
