package models

import "time"

type(
	ShowTimesResource struct {
		Data    []ShowTime	`json:"data"`
	}
	ShowTimeResource struct{
		Data     ShowTime         `json:"data"`
	}
	ShowTime struct {
		Id       	string		`json:"id"`
		Date     	string     	`json:"date"`
		CreatedOn       time.Time       `json:"createdon,omitempty"`
		Movies		[]string	`json:"movies"`
	}
)
