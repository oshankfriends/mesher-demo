package routers

import(
	"github.com/gorilla/mux"
	"github.com/oshankfriends/mesher-demo/booking/controllers"
	"net/http"
)
func InitRoutes()*mux.Router{
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/bookings",controllers.GetBookings).Methods(http.MethodGet)
	router.HandleFunc("/bookings/{userid}",controllers.GetBookingByUserId).Methods(http.MethodGet)
	router.HandleFunc("/bookings",controllers.CreateBooking).Methods(http.MethodPost)
	return router
}
