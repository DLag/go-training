package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email, omitempty"`
}

func main() {
	u := User{"John", "john@example.com"}
	s, _ := json.Marshal(u)
	fmt.Println(string(s))
	s = []byte(`{"name":"Paul"}`)
	fmt.Println(json.Unmarshal(s, &u))
	fmt.Println(u)
}
