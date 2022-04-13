package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func UsersRoute(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Fprint(w, "Users not found")
		return
	}
	fmt.Fprint(w, string(jsonFile))
}

func main() {
	http.HandleFunc("/users", UsersRoute)
	http.ListenAndServe(":8080", nil)
}
