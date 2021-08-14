package main

import (
	"encoding/json"
	"fmt"
)
type person struct {
	First string
	Last  string
	Age   int
}
func main() {
	me := person{
		First: "martin",
		Last:  "cartledge",
		Age:   29,
	}
    bff := person{
		First: "mikel",
		Last:  "howarth",
		Age:   29,
	}
	friends := []person{me, bff}
	fmt.Println(friends)
	// [{martin cartledge 29} {mikel howarth 29}]

	res, err := json.Marshal(friends)

	// Note: json.Marshal converts your JSON to a string of bytes

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	// [91 123 34 70 105 114 115 116 34 58 34 109 97 114 116 105 110 34 44 34 76 97 115 116 34 58 34 99 97 114 116 108 101 100 103 101 34 44 34 65 103 101 34 58 50 57 125 44 123 34 70 105 114 115 116 34 58 34 109 105 107 101 108 34 44 34 76 97 115 116 34 58 34 104 111 119 97 114 116 104 34 44 34 65 103 101 34 58 50 57 125 93]

}
