package q4

import (
	"encoding/json"
	"net/http"
	"time"
)

type Request struct {
	Date string
}
type Response struct {
	Date           string
	LastDayOfMonth int
}

func (r *Request) getResponse() *Response {
	res := &Response{Date: r.Date}
	layout := "2006-01-02T3:04"
	t, err := time.Parse(layout, r.Date)
	if err != nil {
		return res
	}
	year, month, _ := t.Date()
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	year, month, date := lastOfMonth.Date()
	res.LastDayOfMonth = date
	return res
}

func GetLastDayOfMonth(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	res := req.getResponse()
	json.NewEncoder(w).Encode(res)
}
