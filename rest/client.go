package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	resp, err := http.Get("http://localhost:8080/user?id=123")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response from REST server: %+v\n", user)
}
