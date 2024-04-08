package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type product struct {
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Published bool   `json:"published"`
}

func main() {
	jsonData := `{"name": "MacBook Air", "price": 900, "published": true}`

	var p product

	// fazendo a deserialização
	err := json.Unmarshal([]byte(jsonData), &p)

	// verificamos se não teve nenhum erro de parse
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}
