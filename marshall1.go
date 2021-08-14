package main

import (
	"encoding/json"
	"fmt"
)
type person struct {
	Age       int 
	Name      string 
	Address   string
}
type student struct {
       Per_1    person
	   Year_of_study  int
}
   func main() {
	me := person{
		Age: 30,
		Name:  "Gopi",
		Address:   "narasaraopet",
	}
	st := student {
	    Per_1 : me,
		Year_of_study : 2021,
	}
	fmt.Println(st)
	fmt.Println(st.Per_1.Name)
	res, err := json.Marshal(st)

	// Note: json.Marshal converts your JSON to a string of bytes

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("bytestream", res)
	var people []student
	 
	err = json.Unmarshal([]byte(res), &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)
}