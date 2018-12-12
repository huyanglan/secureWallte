package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"wallet/balance"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world", html.EscapeString(r.URL.Path))
}
// Alloccate the address to user
func AllocateAddress(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world", html.EscapeString(r.URL.Path))
}


func QueryBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world", html.EscapeString(r.URL.Path))
	params := mux.Vars(r)
	userInfo := make(chan balance.Account)
	userId := params["userid"]
	currencyId := params["currencyid"]
	go balance.GetBalance(userId, currencyId, userInfo)
	json.NewEncoder(w).Encode(<-userInfo)
}

// withdraw currency to some address
func WithdrawDeal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world", html.EscapeString(r.URL.Path))
}

// query recharge or withdraw status
func QueryStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world", html.EscapeString(r.URL.Path))
}

// query tx list
func QueryTxList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world", html.EscapeString(r.URL.Path))
}

// when it achieve recharge or withdraw, notify the client the status
func NotifyStatus() {
	
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/addressquest/{userid}/{currencyid}/{amount}", AllocateAddress)
	router.HandleFunc("/querybalance/{userid}/{currencyid}", QueryBalance)
	router.HandleFunc("/withdrawrequest", WithdrawDeal)
	router.HandleFunc("/querystatus", QueryStatus)
	router.HandleFunc("/querytxlist", QueryTxList)
	log.Fatal(http.ListenAndServe(":8080", router))
}
