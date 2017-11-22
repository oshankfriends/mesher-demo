package controllers

import (
	"net/http"
	"github.com/oshankfriends/mesher-demo/booking/models"
	"encoding/json"
	"github.com/oshankfriends/mesher-demo/booking/common"
	"github.com/gorilla/mux"
	"errors"
)

var GlobalBookingResource = models.BookingsResource{ Data: make([]models.Booking,0)}

func CreateBooking(w http.ResponseWriter, r *http.Request){
	var dataresource models.BookingResource
	if err := json.NewDecoder(r.Body).Decode(&dataresource); err != nil {
		common.DisplayError(w,err,"Invalid booking data",500)
		return
	}
	GlobalBookingResource.Data = append(GlobalBookingResource.Data,dataresource.Data)
	js,err := json.Marshal(dataresource)
	if err != nil {
		common.DisplayError(w,err,"An unexpected error occured",500)
		return
	}
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func GetBookings(w http.ResponseWriter, r *http.Request){
	js,err := json.Marshal(GlobalBookingResource)
	if err != nil{
		common.DisplayError(w,err,"An unexpected error occured",500)
		return
	}
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func GetBookingByUserId(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userid := vars["userid"]

	for _,booking := range GlobalBookingResource.Data{
		if booking.UserId == userid {
			js,err := json.Marshal(models.BookingResource{Data:booking})
			if err != nil{
				common.DisplayError(w,err,"Unexpected error",500)
				return
			}
			w.Header().Set("Content-Type","application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write(js)
			return
		}
	}
	common.DisplayError(w,errors.New("user not found"),"Failed in getting user information",500)
	return
}