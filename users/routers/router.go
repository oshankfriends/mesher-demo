package routers

import(
	"github.com/gorilla/mux"
	"github.com/oshankfriends/mesher-demo/users/controllers"
	"net/http"
)

func InitRoutes()*mux.Router{
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/users",controllers.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users",controllers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}",controllers.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}/bookings",controllers.GetUserBooking).Methods(http.MethodGet)
	return router
}
