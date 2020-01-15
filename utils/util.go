package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	start      net.IP
	end        net.IP
	ipsTrusted []string
)

func GetValueByEnvKey(envKey string) string {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	return os.Getenv(envKey)
}

func init() {

	fmt.Println(GetValueByEnvKey("db_host"))

	startInit := GetValueByEnvKey("ip_range_start")
	start = net.ParseIP(startInit)

	endInit := GetValueByEnvKey("ip_range_end")
	end = net.ParseIP(endInit)

	ipsTrustedInit := GetValueByEnvKey("ips_trusted")
	s := strings.Split(ipsTrustedInit, ",")
	fmt.Println(s)

	ipsTrusted = s
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckErrDontThrow(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

func containsStringElementInArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CheckIP(ip string) bool {

	//check if ip in our trusted ips range, serve the request if it is
	if GetValueByEnvKey("ip_range_start") == "" || GetValueByEnvKey("ip_range_end") == "" {

		//For target values of Ips
		//Set empty string in .env file in keys ip_range_start or ip_range_end

		if containsStringElementInArray(ipsTrusted, ip) == true {
			fmt.Println("ip IN trusted Ips", ip+"\n", ipsTrusted)
			return true
		} else {
			fmt.Println("ip NOT in trusted Ips", ip+"\n", ipsTrusted)
			return false
		}
	} else {
		//For Ips range from env file
		//Set value string in .env file in keys ip_range_start and ip_range_end

		input := net.ParseIP(ip)
		if input.To4() == nil {
			fmt.Printf("%v is not a valid IPv4 address\n", input)
			return false
		}

		if bytes.Compare(input, start) >= 0 && bytes.Compare(input, end) <= 0 {
			fmt.Printf("%v is between %v and %v\n", input, start, end)
			return true
		}
		fmt.Printf("%v is NOT between %v and %v\n", input, start, end)
		return false
	}
}

func ConvertRequestDateFormatToTimeType(dateFromRequest string) time.Time {

	year, err := strconv.Atoi(strings.SplitN(dateFromRequest, "-", 6)[0])
	CheckErrDontThrow(err)
	log.Println(year)

	month, err := strconv.Atoi(strings.SplitN(dateFromRequest, "-", 6)[1])
	CheckErrDontThrow(err)
	log.Println(month)

	day, err := strconv.Atoi(strings.SplitN(dateFromRequest, "-", 6)[2])
	CheckErrDontThrow(err)
	log.Println(day)

	hours, err := strconv.Atoi(strings.SplitN(dateFromRequest, "-", 6)[3])
	CheckErrDontThrow(err)
	log.Println(hours)

	minutes, err := strconv.Atoi(strings.SplitN(dateFromRequest, "-", 6)[4])
	CheckErrDontThrow(err)
	log.Println(minutes)

	seconds, err := strconv.Atoi(strings.SplitN(dateFromRequest, "-", 6)[5])
	CheckErrDontThrow(err)
	log.Println(seconds)

	date := time.Date(year, time.Month(month), day, hours, minutes, seconds, 0, time.UTC)
	log.Println(date)

	return date
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
