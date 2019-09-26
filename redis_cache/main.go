package main

import (
	"github.com/anderslundholm/go_playground/redis_cache/models"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := models.GetAutocomplete("comp", "ica", 1000)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(res)
	}

	http.HandleFunc("/", indexPage)
	http.ListenAndServe(":3000", nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, http.StatusText(405), 405)
		return
	}
	fmt.Println("hello")
}
