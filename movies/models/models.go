package models

import "time"

type(
	MoviesResource struct {
		Data    []Movie	`json:"data"`
	}
	MovieResource struct{
		Data     Movie         `json:"data"`
	}
	Movie struct {
		Id       	string		`json:"id"`
		Title     	string     	`json:"title"`
		Director        string       	`json:"director"`
		Rating		float32 	`json:"rating"`
		CreatedOn       time.Time	`json:"createdon,omitempty"`
	}
)
