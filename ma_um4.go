package main

import (
	"encoding/json"
	"fmt"
)
type person struct {
	Age       int    `json:"age"`
	Name      string `json:"name"`
	Address   string `json:address"`
}
type student struct {
       Per_1    person `json:"person"`
	   Year_of_study  int  `json:"year_of_study"`
}
func main() {
     m := make(map[string]interface{})
	 m["age"]=20
	 m["name"]="Gopi"
	 m["address"] = "Narasaraopeta"
	 fmt.Println(m)
	
	 m1 := make(map[string]interface{})
	 m1["person"] = m
	 m1["year_of_study"]=4
	// convert map to json
    jsonString, _ := json.Marshal(m1)
    fmt.Println("marshal", string(jsonString))
	 // convert json to struct
    s := student{}
    json.Unmarshal(jsonString, &s)
    fmt.Println("Unamrhal ", s)
	 
}