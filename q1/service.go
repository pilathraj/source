package q1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/pilathraj/talentpro/helper"
)

type Request struct {
	Url string
}

func getWordCount(html string) map[string]int {
	wordcount := make(map[string]int)
	text := strip.StripTags(html)
	words := strings.Fields(text)
	for i := range words {
		wordcount[words[i]]++
	}
	return wordcount
}

func getContent(url string) map[string]int {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return getWordCount(string(html))
}

func GetWordCount(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	res := getContent(req.Url)
	json.NewEncoder(w).Encode(res)
}
