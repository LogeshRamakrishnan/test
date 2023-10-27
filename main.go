package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().String())
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8010", nil))
}

func Test() {
	for {
		fmt.Println(time.Now())
		timer := time.NewTimer(time.Second * 10)
		<-timer.C
		fmt.Println("10 sec")
	}

}
func main() {
	go Test()
	handleRequests()

}
