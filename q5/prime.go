package q5

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type PrimeCounter struct {
	mutex sync.Mutex
	lists []int
}

func (p *PrimeCounter) isPrime(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	isPrime := true
	if value <= 1 {
		isPrime = false
	} else {
		for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
			if value%i == 0 {
				isPrime = false
			}
		}
	}
	if isPrime {
		p.mutex.Lock()
		p.lists = append(p.lists, value)
		p.mutex.Unlock()
	}
}

// getPrime return Prime list by given range
func getPrime(n int) []int {
	wg := new(sync.WaitGroup)
	pc := new(PrimeCounter)
	wg.Add(n + 1)
	for i := 0; i <= n; i++ {
		go pc.isPrime(i, wg)
	}
	wg.Wait()
	return pc.lists
}

func GetPrimeNumbers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	in, _ := strconv.Atoi(params["input"])
	res := getPrime(in)
	json.NewEncoder(w).Encode(res)
}
