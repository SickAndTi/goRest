package controllers

import (
	"fmt"
	"log"
	"net/http"
	"restGoApi/models"
	u "restGoApi/utils"
	"strconv"
	"strings"
	"time"
)

var startTime time.Time
var endTime time.Time
var itemsPerPage int
var currentPageNumber int
var totalItemsCount int

var GetAllClicks = func(w http.ResponseWriter, req *http.Request) {
	log.Println("GET ALL CLICKS")

	clicksFromDb := models.FindAll()
	log.Println(clicksFromDb)

	if clicksFromDb == nil {
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		resp := u.Message(true, "success")
		resp["clicks"] = clicksFromDb
		u.Respond(w, resp)
	}
}

var GetAllClicksPagination = func(w http.ResponseWriter, r *http.Request) {
	log.Println("PAGINATION ALL CLICKS HERE")

	if r.URL.Query().Get("offset") == "" {
		getItemsPerPage, err := strconv.Atoi(u.GetValueByEnvKey("items_per_page"))
		u.CheckErrDontThrow(err)
		fmt.Print("NULL parameter of items per page: ", getItemsPerPage)
		itemsPerPage = getItemsPerPage
	} else {
		getItemsPerPage, err := strconv.Atoi(r.URL.Query().Get("offset"))
		u.CheckErrDontThrow(err)
		fmt.Print("Not null parameter of items per page: ", getItemsPerPage)
		itemsPerPage = getItemsPerPage
	}

	if r.URL.Query().Get("page") == "" {
		currentPageNumber = 1
		fmt.Print("NULL parameter of page number: ", currentPageNumber)
	} else {
		getCurrentPageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))
		u.CheckErrDontThrow(err)
		fmt.Print("Not null parameter page number: ", getCurrentPageNumber)
		currentPageNumber = getCurrentPageNumber
	}

	getTotalItemsCount := models.CountTotal()
	fmt.Print("Total items count from auth: ", getTotalItemsCount)
	totalItemsCount = getTotalItemsCount

	getTotalPageCount := getTotalPageCount(totalItemsCount, itemsPerPage)
	log.Println(getTotalPageCount)

	getCurrentPageClicks := models.GetCurrentPageClicks(currentPageNumber, itemsPerPage)
	fmt.Print("Current page clicks: ", getCurrentPageClicks)

	if getCurrentPageClicks == nil {
		log.Println("PAGINATION FALSE")
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		log.Println("PAGINATION SUCCESS")
		resp := u.Message(true, "success")
		resp["clicks"] = getCurrentPageClicks
		resp["items_per_page"] = itemsPerPage
		resp["current_page_number"] = currentPageNumber
		resp["total_pages"] = getTotalPageCount
		resp["total_items"] = getTotalItemsCount
		u.Respond(w, resp)
	}
}

var GetClickById = func(w http.ResponseWriter, r *http.Request) {
	log.Println("GET_Click By Click ID here")

	clickId := strings.SplitN(r.URL.Path, "/", 5)[4]
	log.Println(clickId)

	clicksFromDb := models.FindOneByClickId(clickId)

	if clicksFromDb == nil {
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		resp := u.Message(true, "success")
		resp["clicks"] = clicksFromDb
		u.Respond(w, resp)
	}
}

var GetAllByOfferId = func(w http.ResponseWriter, r *http.Request) {
	log.Println("GET ALL by OfferId route")

	offerId, err := strconv.Atoi(strings.SplitN(r.URL.Path, "/", 5)[4])
	u.CheckErrDontThrow(err)

	log.Println(offerId)

	clicksFromDb := models.FindAllByOfferId(offerId)
	log.Println(clicksFromDb)

	if clicksFromDb == nil {
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		resp := u.Message(true, "success")
		resp["clicks"] = clicksFromDb
		u.Respond(w, resp)
	}
}

var GetAllByOfferIdPagination = func(w http.ResponseWriter, r *http.Request) {
	log.Println("PAGINATION CLICKS BY offerID HERE")

	offerId, err := strconv.Atoi(strings.SplitN(r.URL.Path, "/", 6)[4])
	u.CheckErrDontThrow(err)

	fmt.Print("OfferId:", offerId)

	if r.URL.Query().Get("offset") == "" {
		getItemsPerPage, err := strconv.Atoi(u.GetValueByEnvKey("items_per_page"))
		u.CheckErrDontThrow(err)
		fmt.Print("NULL parameter of items per page: ", getItemsPerPage)
		itemsPerPage = getItemsPerPage
	} else {
		getItemsPerPage, err := strconv.Atoi(r.URL.Query().Get("offset"))
		u.CheckErrDontThrow(err)
		fmt.Print("Not null parameter of items per page: ", getItemsPerPage)
		itemsPerPage = getItemsPerPage
	}

	if r.URL.Query().Get("page") == "" {
		currentPageNumber = 1
		fmt.Print("NULL parameter of page number: ", currentPageNumber)
	} else {
		getCurrentPageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))
		u.CheckErrDontThrow(err)
		fmt.Print("Not null parameter page number: ", getCurrentPageNumber)
		currentPageNumber = getCurrentPageNumber
	}

	getTotalItemsCountByOfferId := models.CountTotalByOfferId(offerId)
	fmt.Print("Total items count by offerId : ", getTotalItemsCountByOfferId)
	totalItemsCount = getTotalItemsCountByOfferId

	getTotalPageCount := getTotalPageCount(totalItemsCount, itemsPerPage)
	log.Println(getTotalPageCount)

	getCurrentPageClicks := models.FindAllByOfferIdPagination(offerId, currentPageNumber, itemsPerPage)
	fmt.Print("Current page clicks: ", getCurrentPageClicks)

	if getCurrentPageClicks == nil {
		log.Println("PAGINATION FALSE")
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		log.Println("PAGINATION SUCCESS")
		resp := u.Message(true, "success")
		resp["clicks"] = getCurrentPageClicks
		resp["items_per_page"] = itemsPerPage
		resp["current_page_number"] = currentPageNumber
		resp["total_pages"] = getTotalPageCount
		resp["total_items"] = getTotalItemsCountByOfferId
		u.Respond(w, resp)
	}
}

var GetAllByPID = func(w http.ResponseWriter, r *http.Request) {
	log.Println("GET ALL by PID route")

	pid, err := strconv.Atoi(strings.SplitN(r.URL.Path, "/", 5)[4])
	u.CheckErrDontThrow(err)

	log.Println(pid)

	clicksFromDb := models.FindAllByPID(pid)
	log.Println(clicksFromDb)

	if clicksFromDb == nil {
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		resp := u.Message(true, "success")
		resp["clicks"] = clicksFromDb
		u.Respond(w, resp)
	}
}

var GetAllByPIDPagination = func(w http.ResponseWriter, r *http.Request) {
	log.Println("PAGINATION CLICKS BY PID HERE")

	pid, err := strconv.Atoi(strings.SplitN(r.URL.Path, "/", 6)[4])
	u.CheckErrDontThrow(err)

	fmt.Print("PID:", pid)

	if r.URL.Query().Get("offset") == "" {
		getItemsPerPage, err := strconv.Atoi(u.GetValueByEnvKey("items_per_page"))
		u.CheckErrDontThrow(err)
		fmt.Print("NULL parameter of items per page: ", getItemsPerPage)
		itemsPerPage = getItemsPerPage
	} else {
		getItemsPerPage, err := strconv.Atoi(r.URL.Query().Get("offset"))
		u.CheckErrDontThrow(err)
		fmt.Print("Not null parameter of items per page: ", getItemsPerPage)
		itemsPerPage = getItemsPerPage
	}

	if r.URL.Query().Get("page") == "" {
		currentPageNumber = 1
		fmt.Print("NULL parameter of page number: ", currentPageNumber)
	} else {
		getCurrentPageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))
		u.CheckErrDontThrow(err)
		fmt.Print("Not null parameter page number: ", getCurrentPageNumber)
		currentPageNumber = getCurrentPageNumber
	}

	getTotalItemsCountByPID := models.CountTotalByPID(pid)
	fmt.Print("Total items count by PID : ", getTotalItemsCountByPID)
	totalItemsCount = getTotalItemsCountByPID

	getTotalPageCount := getTotalPageCount(totalItemsCount, itemsPerPage)
	log.Println(getTotalPageCount)

	getCurrentPageClicks := models.FindAllByPIDPagination(pid, currentPageNumber, itemsPerPage)
	fmt.Print("Current page clicks: ", getCurrentPageClicks)

	if getCurrentPageClicks == nil {
		log.Println("PAGINATION FALSE")
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		log.Println("PAGINATION SUCCESS")
		resp := u.Message(true, "success")
		resp["clicks"] = getCurrentPageClicks
		resp["items_per_page"] = itemsPerPage
		resp["current_page_number"] = currentPageNumber
		resp["total_pages"] = getTotalPageCount
		resp["total_items"] = getTotalItemsCountByPID
		u.Respond(w, resp)
	}
}

var GetAllByThread = func(w http.ResponseWriter, r *http.Request) {
	log.Println("GET ALL by Thread route")

	thread := strings.SplitN(r.URL.Path, "/", 5)[4]
	log.Println(thread)

	clicksFromDb := models.FindAllByThread(thread)
	log.Println(clicksFromDb)

	if clicksFromDb == nil {
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		resp := u.Message(true, "success")
		resp["clicks"] = clicksFromDb
		u.Respond(w, resp)
	}
}

var GetAllByThreadPagination = func(w http.ResponseWriter, r *http.Request) {
	log.Println("PAGINATION CLICKS BY Thread HERE")

	thread := strings.SplitN(r.URL.Path, "/", 6)[4]

	fmt.Print("Thread:", thread)

	if r.URL.Query().Get("offset") == "" {
		getItemsPerPage, err := strconv.Atoi(u.GetValueByEnvKey("items_per_page"))
		u.CheckErrDontThrow(err)
		fmt.Print("NULL parameter of items per page: ", getItemsPerPage)
		itemsPerPage = getItemsPerPage
	} else {
		getItemsPerPage, err := strconv.Atoi(r.URL.Query().Get("offset"))
		u.CheckErrDontThrow(err)
		fmt.Print("Not null parameter of items per page: ", getItemsPerPage)
		itemsPerPage = getItemsPerPage
	}

	if r.URL.Query().Get("page") == "" {
		currentPageNumber = 1
		fmt.Print("NULL parameter of page number: ", currentPageNumber)
	} else {
		getCurrentPageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))
		u.CheckErrDontThrow(err)
		fmt.Print("Not null parameter page number: ", getCurrentPageNumber)
		currentPageNumber = getCurrentPageNumber
	}

	getTotalItemsCountByThread := models.CountTotalByThread(thread)
	fmt.Print("Total items count by Thread : ", getTotalItemsCountByThread)
	totalItemsCount = getTotalItemsCountByThread

	getTotalPageCount := getTotalPageCount(totalItemsCount, itemsPerPage)
	log.Println(getTotalPageCount)

	getCurrentPageClicks := models.FindAllByThreadPagination(thread, currentPageNumber, itemsPerPage)
	fmt.Print("Current page clicks: ", getCurrentPageClicks)

	if getCurrentPageClicks == nil {
		log.Println("PAGINATION FALSE")
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		log.Println("PAGINATION SUCCESS")
		resp := u.Message(true, "success")
		resp["clicks"] = getCurrentPageClicks
		resp["items_per_page"] = itemsPerPage
		resp["current_page_number"] = currentPageNumber
		resp["total_pages"] = getTotalPageCount
		resp["total_items"] = getTotalItemsCountByThread
		u.Respond(w, resp)
	}
}

var GetAllByPIDDay = func(w http.ResponseWriter, r *http.Request) {
	log.Println("GET ALL by PID time route")

	timeHours, err := strconv.Atoi(strings.SplitN(r.URL.Path, "/", 7)[4])
	u.CheckErrDontThrow(err)
	log.Println(timeHours)

	pid, err := strconv.Atoi(strings.SplitN(r.URL.Path, "/", 7)[6])
	u.CheckErrDontThrow(err)
	log.Println(pid)

	clicksFromDb := models.FindAllByPIDDay(pid, timeHours)
	log.Println(clicksFromDb)

	if clicksFromDb == nil {
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		resp := u.Message(true, "success")
		resp["clicks"] = clicksFromDb
		u.Respond(w, resp)
	}
}

var GetAllByPIDAndSub = func(w http.ResponseWriter, r *http.Request) {
	log.Println("GET ALL by PID and Sub route")

	pid, err := strconv.Atoi(strings.SplitN(r.URL.Path, "/", 5)[3])
	u.CheckErrDontThrow(err)
	log.Println(pid)

	sub := strings.SplitN(r.URL.Path, "/", 5)[4]
	u.CheckErrDontThrow(err)
	log.Println(sub)

	clicksFromDb := models.FindAllByPIDAndSubId(pid, sub)
	log.Println(clicksFromDb)

	if clicksFromDb == nil {
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		resp := u.Message(true, "success")
		resp["clicks"] = clicksFromDb
		u.Respond(w, resp)
	}
}

var GetAllByPIDAndTimeInterval = func(w http.ResponseWriter, r *http.Request) {
	log.Println("GET ALL by PID and Time interval route")

	pid, err := strconv.Atoi(strings.SplitN(r.URL.Path, "/", 9)[4])
	u.CheckErrDontThrow(err)
	log.Println(pid)

	startTimeFromReq := strings.SplitN(r.URL.Path, "/", 9)[6]
	log.Println(startTimeFromReq)

	endTimeFromReq := strings.SplitN(r.URL.Path, "/", 9)[8]
	log.Println(endTimeFromReq)

	startTime = u.ConvertRequestDateFormatToTimeType(startTimeFromReq)
	endTime = u.ConvertRequestDateFormatToTimeType(endTimeFromReq)

	clicksFromDb := models.FindAllByPIDAndTimeInterval(pid, startTime, endTime)
	log.Println(clicksFromDb)

	if clicksFromDb == nil {
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		resp := u.Message(true, "success")
		resp["clicks"] = clicksFromDb
		u.Respond(w, resp)
	}
}

var GetAllByOfferIdAndTimeInterval = func(w http.ResponseWriter, r *http.Request) {
	log.Println("GET ALL by OfferId and Time interval route")

	offerId, err := strconv.Atoi(strings.SplitN(r.URL.Path, "/", 9)[4])
	u.CheckErrDontThrow(err)
	log.Println(offerId)

	startTimeFromReq := strings.SplitN(r.URL.Path, "/", 9)[6]
	log.Println(startTimeFromReq)

	endTimeFromReq := strings.SplitN(r.URL.Path, "/", 9)[8]
	log.Println(endTimeFromReq)

	startTime = u.ConvertRequestDateFormatToTimeType(startTimeFromReq)
	endTime = u.ConvertRequestDateFormatToTimeType(endTimeFromReq)

	clicksFromDb := models.FindAllByOfferIdAndTimeInterval(offerId, startTime, endTime)
	log.Println(clicksFromDb)

	if clicksFromDb == nil {
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		resp := u.Message(true, "success")
		resp["clicks"] = clicksFromDb
		u.Respond(w, resp)
	}
}

var GetAllByThreadAndTimeInterval = func(w http.ResponseWriter, r *http.Request) {
	log.Println("GET ALL by Thread and Time interval route")

	thread := strings.SplitN(r.URL.Path, "/", 9)[4]
	log.Println(thread)

	startTimeFromReq := strings.SplitN(r.URL.Path, "/", 9)[6]
	log.Println(startTimeFromReq)

	endTimeFromReq := strings.SplitN(r.URL.Path, "/", 9)[8]
	log.Println(endTimeFromReq)

	startTime = u.ConvertRequestDateFormatToTimeType(startTimeFromReq)
	endTime = u.ConvertRequestDateFormatToTimeType(endTimeFromReq)

	clicksFromDb := models.FindAllByThreadAndTimeInterval(thread, startTime, endTime)
	log.Println(clicksFromDb)

	if clicksFromDb == nil {
		resp := u.Message(false, "No data with your parameters")
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, resp)
	} else {
		resp := u.Message(true, "success")
		resp["clicks"] = clicksFromDb
		u.Respond(w, resp)
	}
}

func getTotalPageCount(totalItemsCount int, itemsPerPage int) int {
	var countPagesTotal int

	fmt.Print("ITEMS PER PAGE :", itemsPerPage)
	if totalItemsCount%itemsPerPage == 0 {
		countPagesTotal = totalItemsCount / itemsPerPage
	} else {
		countPagesTotal = totalItemsCount/itemsPerPage + 1
	}
	return countPagesTotal
}
