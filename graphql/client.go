package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	query := `{"query": "{ user(id:\"123\"){ id name } }"}`
	resp, err := http.Post("http://localhost:8081/graphql", "application/json", bytes.NewBuffer([]byte(query)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response from GraphQL server: %+v\n", result)
}
