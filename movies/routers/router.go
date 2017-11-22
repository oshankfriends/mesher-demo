package routers

import(
	"github.com/gorilla/mux"
	"github.com/oshankfriends/mesher-demo/movies/controllers"
	"net/http"
)
func InitRoutes()*mux.Router{
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/movies",controllers.GetMovies).Methods(http.MethodGet)
	router.HandleFunc("/movies",controllers.CreateMovie).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}",controllers.GetMovieById).Methods(http.MethodGet)
	router.HandleFunc("/movies/{id}",controllers.DeleteMovie).Methods(http.MethodDelete)
	return router
}
