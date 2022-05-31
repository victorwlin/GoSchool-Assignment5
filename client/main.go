package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	r := mux.NewRouter()

	r.Handle("/favicon.ico", http.NotFoundHandler())

	r.HandleFunc("/", index)
	r.HandleFunc("/friend/", showFriend)
	r.HandleFunc("/addfriend/", addFriend)
	r.HandleFunc("/editfriend/", editFriend)
	r.HandleFunc("/deletefriend/", deleteFriend)

	http.Handle("/", r)

	http.ListenAndServe(":5221", nil)
}
