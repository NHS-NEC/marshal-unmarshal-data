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
   func main() {
     m := make(map[string]interface{})
	 m["age"]=20
	 m["name"]="Gopi"
	 m["address"] = "Narasaraopeta"
	 fmt.Println(m)
	
	// convert map to json
    jsonString, _ := json.Marshal(m)
    fmt.Println(string(jsonString))
	 // convert json to struct
    s := person{}
    json.Unmarshal(jsonString, &s)
    fmt.Println(s)
	 
}