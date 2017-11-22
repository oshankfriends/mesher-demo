package models

import "time"

type (
	UsersResource struct {
		Data []User `json:"data"`
	}
	UserResource struct {
		Data User `json:"data"`
	}
	User struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		LastName string `json:"lastname"`
	}
)

type (
	BookingResult struct {
		Data Booking `json:"data"`
	}

	Booking struct {
		UserId     string   `json:"userid"`
		ShowTimeId string   `json:"showtimeid"`
		Movies     []string `json:"movies"`
	}
)
type (
	ShowTimeResult struct {
		Data ShowTime `json:"data"`
	}
	ShowTime struct {
		Id        string    `json:"id"`
		Date      string    `json:"date"`
		CreatedOn time.Time `json:"createdon,omitempty"`
		Movies    []string  `json:"movies"`
	}
)
type (
	MovieResult struct {
		Data Movie `json:"data"`
	}
	Movie struct {
		Id        string    `json:"id"`
		Title     string    `json:"title"`
		Director  string    `json:"director"`
		Rating    float32   `json:"rating"`
		CreatedOn time.Time `json:"createdon,omitempty"`
	}
)

type Result struct {
	User     User     `json:"user"`
	Booking  Booking  `json:"booking"`
	Showtime ShowTime `json:"showtime"`
	Movies   []Movie  `json:"movies"`
}
