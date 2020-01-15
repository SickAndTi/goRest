package models

import (
	"fmt"
	"log"
	u "restGoApi/utils"
	"time"
)

type ClickModel struct {
	ClickID         string    `db:"click_id"          json:"click_id"`
	PreviousClickID string    `db:"previous_click_id" json:"previous_click_id"`
	Thread          string    `db:"thread"            json:"thread"`
	OfferId         int       `db:"offer_id"          json:"offer_id"`
	PID             int       `db:"pid"               json:"pid"`
	Headers         string    `db:"headers"           json:"headers"`
	UserAgent       string    `db:"user_agent"        json:"user_agent"`
	Cookie          string    `db:"cookie"            json:"cookie"`
	Host            string    `db:"host"              json:"host"`
	RemoteAddr      string    `db:"remote_addr"       json:"remote_addr"`
	UtmSource       string    `db:"utm_source"        json:"utm_source"`
	UtmCampaign     string    `db:"utm_campaign"      json:"utm_campaign"`
	UtmMedium       string    `db:"utm_medium"        json:"utm_medium"`
	UtmContent      string    `db:"utm_content"       json:"utm_content"`
	SubId1          string    `db:"sub_id_1"          json:"sub_id_1"`
	SubId2          string    `db:"sub_id_2"          json:"sub_id_2"`
	SubId3          string    `db:"sub_id_3"          json:"sub_id_3"`
	SubId4          string    `db:"sub_id_4"          json:"sub_id_4"`
	SubId5          string    `db:"sub_id_5"          json:"sub_id_5"`
	SubId6          string    `db:"sub_id_6"          json:"sub_id_6"`
	Date            time.Time `db:"date"              json:"date"`
	DateUnix        int64     `db:"date_unix"         json:"date_unix"`
}

func (click *ClickModel) Validate() (map[string]interface{}, bool) {

	//click must be unique
	temp := &ClickModel{}

	if temp.ClickID == "" {
		return u.Message(false, "NO click id!!!"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func FindAll() []ClickModel {

	log.Println("FIND ALL")

	var c []ClickModel

	if err := u.GetDB().Select(&c, "SELECT click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks"); err != nil {
		u.CheckErrDontThrow(err)
	}
	return c
}

func FindOneByClickId(clickId string) []ClickModel {

	var c []ClickModel
	if err := u.GetDB().Select(&c, "SELECT click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE click_id=?", clickId); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}

func FindAllByOfferId(offerId int) []ClickModel {

	var c []ClickModel
	if err := u.GetDB().Select(&c, "SELECT click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE offer_id=?", offerId); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}

func FindAllByOfferIdPagination(offerId int, currentPageNumberFromController int, itemsPerPageFromController int) []ClickModel {

	var c []ClickModel

	offset := (currentPageNumberFromController - 1) * itemsPerPageFromController

	fmt.Print("OFFSET : %a", offset)

	if err := u.GetDB().Select(&c, "SELECT click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE offer_id=? ORDER BY date LIMIT ?, ?  ", offerId, offset, itemsPerPageFromController); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}

func FindAllByPID(pid int) []ClickModel {

	var c []ClickModel
	if err := u.GetDB().Select(&c, "SELECT	click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE pid=?", pid); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}

func FindAllByPIDPagination(pid int, currentPageNumberFromController int, itemsPerPageFromController int) []ClickModel {
	var c []ClickModel
	offset := (currentPageNumberFromController - 1) * itemsPerPageFromController

	fmt.Print("OFFSET : %a", offset)

	if err := u.GetDB().Select(&c, "SELECT click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE pid=? ORDER BY date LIMIT ?, ?  ", pid, offset, itemsPerPageFromController); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}

func FindAllByThread(thread string) []ClickModel {

	var clicks []ClickModel

	if err := u.GetDB().Select(&clicks, "SELECT click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE thread=?", thread); err != nil {
		u.CheckErrDontThrow(err)
	}

	return clicks
}

func FindAllByThreadPagination(thread string, currentPageNumberFromController int, itemsPerPageFromController int) []ClickModel {
	var c []ClickModel
	offset := (currentPageNumberFromController - 1) * itemsPerPageFromController

	fmt.Print("OFFSET : %a", offset)

	if err := u.GetDB().Select(&c, "SELECT click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE thread=? ORDER BY date LIMIT ?, ?  ", thread, offset, itemsPerPageFromController); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}

func FindAllByPIDDay(pid int, lastTimeHours int) []ClickModel {

	var c []ClickModel
	if err := u.GetDB().Select(&c, "SELECT	click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE pid=? AND date >= (now() - toIntervalHour(?)) ", pid, lastTimeHours); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}

func FindAllByPIDAndSubId(pid int, sub string) []ClickModel {

	var c []ClickModel
	if err := u.GetDB().Select(&c, "SELECT	click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE pid=? AND sub_id_1=?", pid, sub); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}

func FindAllByPIDAndTimeInterval(pid int, startTimeInterval time.Time, endTimeInterval time.Time) []ClickModel {

	var c []ClickModel
	if err := u.GetDB().Select(&c, "SELECT click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE pid=? AND date BETWEEN ? AND ? ", pid, startTimeInterval, endTimeInterval); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}

func FindAllByOfferIdAndTimeInterval(offerId int, startTimeInterval time.Time, endTimeInterval time.Time) []ClickModel {

	var c []ClickModel
	if err := u.GetDB().Select(&c, "SELECT click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE offer_id=? AND date BETWEEN ? AND ? ", offerId, startTimeInterval, endTimeInterval); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}

func FindAllByThreadAndTimeInterval(thread string, startTimeInterval time.Time, endTimeInterval time.Time) []ClickModel {

	var c []ClickModel
	if err := u.GetDB().Select(&c, "SELECT click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks WHERE thread=? AND date BETWEEN ? AND ? ", thread, startTimeInterval, endTimeInterval); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}

func CountTotal() int {
	var countItemsTotal int

	if err := u.GetDB().Get(&countItemsTotal, "SELECT count(*) FROM clicks"); err != nil {
		u.CheckErrDontThrow(err)
	}

	log.Println("COUNT total items from model")
	log.Println(countItemsTotal)
	return countItemsTotal
}

func CountTotalByOfferId(offerId int) int {
	var countItemsTotal int

	if err := u.GetDB().Get(&countItemsTotal, "SELECT count(*) FROM clicks WHERE offer_id = ?", offerId); err != nil {
		u.CheckErrDontThrow(err)
	}

	fmt.Print("COUNT total items from model by offerId: ", countItemsTotal)
	return countItemsTotal
}

func CountTotalByPID(pid int) int {
	var countItemsTotal int

	if err := u.GetDB().Get(&countItemsTotal, "SELECT count(*) FROM clicks WHERE pid = ?", pid); err != nil {
		u.CheckErrDontThrow(err)
	}

	fmt.Print("COUNT total items from model by pid: ", countItemsTotal)
	return countItemsTotal
}

func CountTotalByThread(thread string) int {

	var countItemsTotal int

	if err := u.GetDB().Get(&countItemsTotal, "SELECT count(*) FROM clicks WHERE thread = ?", thread); err != nil {
		u.CheckErrDontThrow(err)
	}

	fmt.Print("COUNT total items from model by thread: ", countItemsTotal)
	return countItemsTotal
}

func GetCurrentPageClicks(currentPageNumberFromController int, itemsPerPageFromController int) []ClickModel {

	var c []ClickModel

	offset := (currentPageNumberFromController - 1) * itemsPerPageFromController

	fmt.Print("OFFSET : %a", offset)

	if err := u.GetDB().Select(&c, "SELECT click_id, previous_click_id, thread, offer_id, pid, headers, user_agent, cookie, host, remote_addr, utm_source, utm_campaign, utm_medium, utm_content, sub_id_1, sub_id_2, sub_id_3, sub_id_4, sub_id_5, sub_id_6, date, date_unix FROM clicks ORDER BY date LIMIT ?, ?  ", offset, itemsPerPageFromController); err != nil {
		u.CheckErrDontThrow(err)
	}

	return c
}
