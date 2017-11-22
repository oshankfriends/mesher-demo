package controllers

import (
	"net/http"
	"github.com/oshankfriends/mesher-demo/showtimes/models"
	"encoding/json"
	"github.com/oshankfriends/mesher-demo/showtimes/common"
	"github.com/gorilla/mux"
	"errors"
)

var GlobalShowTimeResource = models.ShowTimesResource{ Data: make([]models.ShowTime,0)}


func CreateShowTime(w http.ResponseWriter, r *http.Request){
	var dataresource models.ShowTimeResource
	if err := json.NewDecoder(r.Body).Decode(&dataresource); err != nil {
		common.DisplayError(w,err,"Invalid showtime data",500)
		return
	}
	GlobalShowTimeResource.Data = append(GlobalShowTimeResource.Data,dataresource.Data)
	js,err := json.Marshal(dataresource)
	if err != nil {
		common.DisplayError(w,err,"An unexpected error occured",500)
		return
	}
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func GetShowTimes(w http.ResponseWriter, r *http.Request){
	js,err := json.Marshal(GlobalShowTimeResource)
	if err != nil{
		common.DisplayError(w,err,"An unexpected error occured",500)
		return
	}
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func GetShowTimeByDate(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	date := vars["date"]

	for _,showTime := range GlobalShowTimeResource.Data{
		if showTime.Date == date {
			js,err := json.Marshal(models.ShowTimeResource{Data:showTime})
			if err != nil {
				common.DisplayError(w,err,"An unexpected error occured",500)
				return
			}
			w.Header().Set("Content-Type","application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write(js)
			return
		}
	}
	common.DisplayError(w,errors.New("error in getting show"),"show not exist",500)
}

func DeleteShowTime(w http.ResponseWriter, r *http.Request){
	pathParameters := mux.Vars(r)
	id := pathParameters["id"]

	for index,showtime := range GlobalShowTimeResource.Data{
		if showtime.Id == id {
			GlobalShowTimeResource.Data = append(GlobalShowTimeResource.Data[:index],GlobalShowTimeResource.Data[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	common.DisplayError(w,errors.New("error in getting show"),"Show not present with given showId",500)
}

func GetShowTimeById(w http.ResponseWriter, r *http.Request){
	pathParameters := mux.Vars(r)
	id := pathParameters["id"]

	for _,showtime := range GlobalShowTimeResource.Data{
		if showtime.Id == id {
			js,err := json.Marshal(models.ShowTimeResource{Data:showtime})
			if err != nil {
				common.DisplayError(w,err,"An unexpected error occured",500)
				return
			}
			w.Header().Set("Content-Type","application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write(js)
			return
		}
	}
	common.DisplayError(w,errors.New("error in getting show"),"show not exist",500)
}
