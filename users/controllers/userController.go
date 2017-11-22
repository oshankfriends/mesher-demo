package controllers

import (
	"net/http"
	"github.com/oshankfriends/mesher-demo/users/models"
	"encoding/json"
	"github.com/oshankfriends/mesher-demo/users/common"
	"github.com/gorilla/mux"
	"context"
	"time"
	"fmt"
	"errors"
)

var GlobalUserResource = models.UsersResource{ Data: make([]models.User,0)}

func CreateUser(w http.ResponseWriter, r *http.Request){
	var dataresource models.UserResource
	if err := json.NewDecoder(r.Body).Decode(&dataresource); err != nil {
		common.DisplayError(w,err,"Invalid user data",500)
		return
	}
	GlobalUserResource.Data = append(GlobalUserResource.Data,dataresource.Data)
	js,err := json.Marshal(dataresource)
	if err != nil {
		common.DisplayError(w,err,"An unexpected error occured",500)
		return
	}
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func GetUsers(w http.ResponseWriter, r *http.Request){
	js,err := json.Marshal(GlobalUserResource)
	if err != nil{
		common.DisplayError(w,err,"An unexpected error occured",500)
		return
	}
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	pathParameters := mux.Vars(r)
	id := pathParameters["id"]

	for index,user := range GlobalUserResource.Data{
		if user.Id == id {
			GlobalUserResource.Data = append(GlobalUserResource.Data[:index],GlobalUserResource.Data[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	common.DisplayError(w,errors.New("error in getting movie"),"movie not present with given id",500)
}

func GetUserBooking(w http.ResponseWriter, r *http.Request){
	var(
		ctx context.Context
		cancel context.CancelFunc
		result = models.Result{}
		movies = make([]models.Movie,0)
	)

	timeout,err := time.ParseDuration(r.FormValue("timeout"))
	if err == nil {
		ctx,cancel = context.WithTimeout(context.Background(),timeout)
	} else {
		ctx,cancel = context.WithCancel(context.Background())
	}
	defer cancel()
	vars := mux.Vars(r)
	userid := vars["id"]
	bookingresult,err := GetBookingDetails(ctx,userid)
	if err != nil {
		common.DisplayError(w,err,"Booking Details not found",500)
		return
	}
	showTimeId := bookingresult.Data.ShowTimeId
	showTimeresult,err := GetShowTimeDetails(ctx,showTimeId)
	if err != nil {
		common.DisplayError(w,err,"ShowTime Details not found",500)
		return
	}

	moviesId := bookingresult.Data.Movies
	for _,movieId := range moviesId{
		moviesresult,err := GetMoviesDetails(ctx,movieId)
		if err == nil {
			movies = append(movies,moviesresult.Data)
		}
	}
	result.Booking = bookingresult.Data
	result.Showtime = showTimeresult.Data
	result.User = GetUserDetails(userid)
	result.Movies = movies

	js,err := json.Marshal(result)
	if err != nil{
		common.DisplayError(w,err,"An unexpected error occured",500)
		return
	}
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func GetUserDetails(userid string)models.User{
	for _,user := range GlobalUserResource.Data{
		if userid == user.Id{
			return user
		}
	}
	return models.User{}
}

func GetBookingDetails(ctx context.Context, userid string)(models.BookingResult,error){
	url := fmt.Sprintf("http://%s/bookings/%s",common.AppConfig.BookingServer,userid)
	req,err := http.NewRequest(http.MethodGet,url,nil)
	if err != nil{
		return models.BookingResult{},err
	}
	var bookingResult models.BookingResult
	err = httpDo(ctx,req,func(resp *http.Response,err error)error{
		if err != nil {
			return err
		} else if resp.StatusCode != http.StatusOK {
			return errors.New(fmt.Sprintf("Recieved %d status code from bookingservice",resp.StatusCode))
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&bookingResult); err != nil {
			return err
		}
		return nil
	})
	return bookingResult,err
}

func GetShowTimeDetails(ctx context.Context, showtimeId string)(models.ShowTimeResult,error){
	url := fmt.Sprintf("http://%s/showtimes/%s",common.AppConfig.ShowTimeServer,showtimeId)
	req,err := http.NewRequest(http.MethodGet,url,nil)
	if err != nil{
		return models.ShowTimeResult{},err
	}
	var showTimeResult models.ShowTimeResult
	err = httpDo(ctx,req,func(resp *http.Response,err error)error{
		if err != nil {
			return err
		} else if resp.StatusCode != http.StatusOK {
			return errors.New(fmt.Sprintf("Recieved %d status code from showtime",resp.StatusCode))
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&showTimeResult); err != nil {
			return err
		}
		return nil
	})
	return showTimeResult,err
}

func GetMoviesDetails(ctx context.Context, movieId string)(models.MovieResult,error){
	url := fmt.Sprintf("http://%s/movies/%s",common.AppConfig.MovieServer,movieId)
	req,err := http.NewRequest(http.MethodGet,url,nil)
	if err != nil{
		return models.MovieResult{},err
	}
	var movieResult models.MovieResult
	err = httpDo(ctx,req,func(resp *http.Response,err error)error{
		if err != nil {
			return err
		} else if resp.StatusCode != http.StatusOK {
			return errors.New(fmt.Sprintf("Recieved %d status code from showtime",resp.StatusCode))
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&movieResult); err != nil {
			return err
		}
		return nil
	})
	return movieResult,err
}


func httpDo(ctx context.Context,req *http.Request, callBack func(*http.Response,error)error) error {
	tr := &http.Transport{}
	client := http.Client{Transport: tr}
	errChan := make(chan error,1)
	go func() {errChan <- callBack(client.Do(req))}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-errChan
		return ctx.Err()
	case err := <- errChan:
		return err

	}
}
