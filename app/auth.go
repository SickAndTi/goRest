package app

import (
	"log"
	"net/http"
	"os"
	u "restGoApi/utils"
	"strings"
)

var endpointsNoAuth []string

func init() {
	endpointsNoAuthInit := u.GetValueByEnvKey("endpointsNoAuth")
	s := strings.Split(endpointsNoAuthInit, ",")
	endpointsNoAuth = s
	log.Println(s)
}

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestPath := r.URL.Path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, endpoint := range endpointsNoAuth {
			log.Println("Endpoint: " + endpoint)
			log.Println("RequestPath: " + requestPath)
			if endpoint == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		log.Println(splitted)

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in

		log.Println(tokenPart)

		if tokenPart == os.Getenv("token") {
			next.ServeHTTP(w, r) //proceed in the middleware chain!
		} else {
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
	})
}

var ProxyIps = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userIP := strings.Split(r.RemoteAddr, ":")[0]
		log.Println("UserIp: " + userIP)
		if u.CheckIP(userIP) == true {
			next.ServeHTTP(w, r)
			return
		} else {
			resp := u.Message(false, "No data for your IP")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, resp)
			return
		}
	})
}

//var PaginationResponse = func(next http.Handler) http.Handler {
//
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

//		log.Println("PAGINATION HERE")
//		var itemsPerPage int
//		var currentPageNumber int
//
//		if r.URL.Query().Get("offset") == "" {
//			itemsPerPage,err  := strconv.Atoi(u.GetValueByEnvKey("items_per_page"))
//			u.CheckErr(err)
//			fmt.Print("NULL parameter of items per page: %a" , itemsPerPage)
//		} else {
//			getItemsPerPage,err := strconv.Atoi(r.URL.Query().Get("offset"))
//			u.CheckErr(err)
//			fmt.Print("Not null parameter of items per page: %a ", getItemsPerPage)
//			itemsPerPage = getItemsPerPage
//		}
//
//		if r.URL.Query().Get("page") == "" {
//			currentPageNumber = 1
//			fmt.Print("NULL parameter of page number: %a " , currentPageNumber)
//		} else {
//			getCurrentPageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))
//			u.CheckErr(err)
//			fmt.Print("Not null parameter page number: %a" , getCurrentPageNumber)
//			currentPageNumber = getCurrentPageNumber
//		}
//
//		getTotalItemsCount := models.CountTotal()
//		fmt.Print("Total items count from auth: %a", getTotalItemsCount)
//		totalItemsCount := getTotalItemsCount
//
//		getTotalPageCount := getTotalPageCount(totalItemsCount, itemsPerPage)
//		log.Println(getTotalPageCount)
//
//		getCurrentPageClicks := models.GetCurrentPageClicks(currentPageNumber, itemsPerPage)
//		fmt.Print("Current page clicks: %a" ,getCurrentPageClicks)
//
//		if getCurrentPageClicks == nil {
//			log.Println("PAGINATION FALSE")
//			resp := u.Message(false, "No data with your parameters")
//			w.WriteHeader(http.StatusNotFound)
//			u.Respond(w, resp)
//			return
//		} else {
//			log.Println("PAGINATION SUCCESS")
//			//resp := u.Message(true, "success")
//			//resp["items_per_page"] = itemsPerPage
//			//resp["current_page_number"] = currentPageNumber
//			//resp["total_pages"] = getTotalPageCount
//			//resp["total_items"] = getTotalItemsCount
//			next.ServeHTTP(w, r)
//			return
//		}
//	})
//}
//
//func getTotalPageCount(totalItemsCount int, itemsPerPage int) int {
//	var countPagesTotal int
//
//	if totalItemsCount%itemsPerPage == 0 {
//		countPagesTotal = totalItemsCount / itemsPerPage
//	} else {
//		countPagesTotal = totalItemsCount/itemsPerPage + 1
//	}
//	return countPagesTotal
//}
