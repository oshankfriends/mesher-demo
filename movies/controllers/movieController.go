package controllers

import (
	"net/http"
	"github.com/oshankfriends/mesher-demo/movies/models"
	"encoding/json"
	"github.com/oshankfriends/mesher-demo/movies/common"
	"github.com/gorilla/mux"
	"errors"
)

var GlobalMovieResource = models.MoviesResource{ Data: make([]models.Movie,0)}


func CreateMovie(w http.ResponseWriter, r *http.Request){
	var dataresource models.MovieResource
	if err := json.NewDecoder(r.Body).Decode(&dataresource); err != nil {
		common.DisplayError(w,err,"Invalid movie data",500)
		return
	}
	GlobalMovieResource.Data = append(GlobalMovieResource.Data,dataresource.Data)
	js,err := json.Marshal(dataresource)
	if err != nil {
		common.DisplayError(w,err,"An unexpected error occured",500)
		return
	}
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func GetMovies(w http.ResponseWriter, r *http.Request){
	js,err := json.Marshal(GlobalMovieResource)
	if err != nil{
		common.DisplayError(w,err,"An unexpected error occured",500)
		return
	}
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func GetMovieById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	for _,movie := range GlobalMovieResource.Data{
		if movie.Id == id {
			js,err := json.Marshal(models.MovieResource{Data:movie})
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
	common.DisplayError(w,errors.New("error in getting movie"),"movie not exist",500)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	pathParameters := mux.Vars(r)
	id := pathParameters["id"]

	for index,movie := range GlobalMovieResource.Data{
		if movie.Id == id {
			GlobalMovieResource.Data = append(GlobalMovieResource.Data[:index],GlobalMovieResource.Data[index+1:]...)
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

