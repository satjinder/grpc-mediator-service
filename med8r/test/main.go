package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{
        "fruits" : {
            "a": "apple",
            "b": "banana"
        },
        "colors" : {
            "r": "red",
            "g": "green"
        }
    }`

	var x map[string]interface{}

	json.Unmarshal([]byte(jsonStr), &x)
	fmt.Println(x["fruits"])
}
