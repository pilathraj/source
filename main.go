package main

import (
	"log"
	"net/http"

	"github.com/pilathraj/talentpro/q1"
	"github.com/pilathraj/talentpro/q2"
	"github.com/pilathraj/talentpro/q3"
	"github.com/pilathraj/talentpro/q4"
	"github.com/pilathraj/talentpro/q5"

	"github.com/gorilla/mux"
)

func server() {
	r := mux.NewRouter()

	// Handle all preflight request
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	r.HandleFunc("/q1", q1.GetWordCount).Methods("POST")
	r.HandleFunc("/q2", q2.GetExcelInfo).Methods("POST")

	r.HandleFunc("/q3/user", q3.GetUsers).Methods("GET")
	r.HandleFunc("/q3/user/{id}", q3.GetUser).Methods("GET")
	r.HandleFunc("/q3/user", q3.CreateUser).Methods("POST")
	r.HandleFunc("/q3/user/{id}", q3.UpdateUser).Methods("PUT")
	r.HandleFunc("/q3/user/{id}", q3.DeleteUser).Methods("DELETE")

	r.HandleFunc("/q4", q4.GetLastDayOfMonth).Methods("POST")
	r.HandleFunc("/q5/{input}", q5.GetPrimeNumbers).Methods("GET")

	//r.Use(mux.CORSMethodMiddleware(r))

	log.Fatal(http.ListenAndServe(":5000", r))
}

func main() {
	server()
}
