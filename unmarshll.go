package main

import (
	"encoding/json"
	"fmt"
)
type person struct {
	First string
	Last string
	Age int
}
func main() {
	rawData := `[{"First":"martin","Last":"cartledge","Age":29},{"First":"mikel","Last":"howarth","Age":29}]`
	fmt.Println(rawData)
	byteString := []byte(rawData)
	fmt.Println(byteString)
	// people := []person{}
	var people []person
	err := json.Unmarshal(byteString, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
	for i, v := range people {
		fmt.Println("index: ", i)
		fmt.Println("First: ", v.First, "Last: ", v.Last, "Age: ", v.Age)
	}
}