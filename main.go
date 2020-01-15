package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restGoApi/app"
	"restGoApi/controllers"
	u "restGoApi/utils"
)

func main() {

	log.Println("HI THERE main.go")

	router := mux.NewRouter()

	router.Use(app.JwtAuthentication) //attach JWT auth middleware
	router.Use(app.ProxyIps)          //attach correct IPs
	//router.Use(app.PaginationResponse) //attach pagination in response

	router.HandleFunc("/api/clicks", controllers.GetAllClicks).Methods("GET")
	router.HandleFunc("/api/clicks/pagination", controllers.GetAllClicksPagination).Methods("GET")
	router.HandleFunc("/api/click/id/{id}", controllers.GetClickById).Methods("GET")
	router.HandleFunc("/api/click/offer/{offerId}", controllers.GetAllByOfferId).Methods("GET")
	router.HandleFunc("/api/click/offer/{offerId}/pagination", controllers.GetAllByOfferIdPagination).Methods("GET")
	router.HandleFunc("/api/click/thread/{thread}", controllers.GetAllByThread).Methods("GET")
	router.HandleFunc("/api/click/thread/{thread}/pagination", controllers.GetAllByThreadPagination).Methods("GET")
	router.HandleFunc("/api/click/pid/{pid}", controllers.GetAllByPID).Methods("GET")
	router.HandleFunc("/api/click/pid/{pid}/pagination", controllers.GetAllByPIDPagination).Methods("GET")
	router.HandleFunc("/api/click/date/{time}/pid/{pid}", controllers.GetAllByPIDDay).Methods("GET")
	router.HandleFunc("/api/sub/{pid}/{sub}", controllers.GetAllByPIDAndSub).Methods("GET")
	router.HandleFunc("/api/clicksByPidAndTime/pid/{pid}/start/{start}/end/{end}", controllers.GetAllByPIDAndTimeInterval).Methods("GET")
	router.HandleFunc("/api/clicksByOfferIdAndTime/offer/{offerId}/start/{start}/end/{end}", controllers.GetAllByOfferIdAndTimeInterval).Methods("GET")
	router.HandleFunc("/api/clicksByThreadAndTime/thread/{thread}/start/{start}/end/{end}", controllers.GetAllByThreadAndTimeInterval).Methods("GET")

	port := u.GetValueByEnvKey("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
