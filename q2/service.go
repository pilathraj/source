package q2

import (
	"encoding/json"
	"net/http"

	"github.com/pilathraj/talentpro/helper"
)

const LetterCount = 26

type Request struct {
	ColumnStart string
	Rows        int
	Columns     int
}
type Response struct {
	Data interface{}
}

func lettersToInt(s string) int {
	n := 0
	for _, c := range s {
		n = n*LetterCount + 1 + int(c) - int('A')
	}
	return n
}

func intToLetters(number int32) string {
	letters := ""
	number--
	if firstLetter := number / LetterCount; firstLetter > 0 {
		letters += intToLetters(firstLetter)
		letters += string('A' + number%LetterCount)
	} else {
		letters += string('A' + number)
	}
	return letters
}

func (r *Request) getExcelInfo() []string {
	var res []string
	size := r.Rows * r.Columns
	startPos := lettersToInt(r.ColumnStart)
	endPos := startPos + size
	for i := startPos; i < endPos; i++ {
		v := intToLetters(int32(i))
		res = append(res, v)
	}
	return res

}

func GetExcelInfo(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rows := req.getExcelInfo()
	data := make([][]string, req.Rows)
	for i := 0; i < req.Rows; i++ {
		data[i] = rows[i*req.Columns : (i+1)*req.Columns]
	}
	json.NewEncoder(w).Encode(Response{Data: data})
}
