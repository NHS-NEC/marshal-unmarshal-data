package main

import (
    "encoding/base64"
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

    // String to encode
    //str := "Hello World!"
	 
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
    // base64.StdEncoding: Standard encoding with padding
    // It requires a byte slice so we cast the string to []byte
    encodedStr := base64.StdEncoding.EncodeToString([]byte(jsonString))
    fmt.Println("Encoded:", encodedStr)

    // Decoding may return an error, in case the input is not well formed
    decodedStr, err := base64.StdEncoding.DecodeString(encodedStr)
    if err != nil {
        panic("malformed input")
    }
    fmt.Println("Decoded:", string(decodedStr))
	// convert json to struct
    s := student{}
    json.Unmarshal(jsonString, &s)
    fmt.Println("Unamrhal ", s)
}