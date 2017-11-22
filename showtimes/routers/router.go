package routers

import(
	"github.com/gorilla/mux"
	"github.com/oshankfriends/mesher-demo/showtimes/controllers"
	"net/http"
)
func InitRoutes()*mux.Router{
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/showtimes",controllers.GetShowTimes).Methods(http.MethodGet)
//	router.HandleFunc("/showtimes/{date}",controllers.GetShowTimeByDate).Methods(http.MethodGet)
	router.HandleFunc("/showtimes",controllers.CreateShowTime).Methods(http.MethodPost)
	router.HandleFunc("/showtimes/{id}",controllers.DeleteShowTime).Methods(http.MethodDelete)
	router.HandleFunc("/showtimes/{id}",controllers.GetShowTimeById).Methods(http.MethodGet)
	return router
}
