package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	Name string `json:"name"`
}

func main() {
	stream := `
	{"name": "Fred"}
	{"name": "Mary"}
	{"a": "b"}
	{}
	`

	dec := json.NewDecoder(strings.NewReader(stream))
	for dec.More() {
		var p person
		err := dec.Decode(&p)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", p)
	}

}
