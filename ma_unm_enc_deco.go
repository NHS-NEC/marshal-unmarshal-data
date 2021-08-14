package main

import (
	"encoding/json"
	"encoding/base64"
	"fmt"
)
type re_person struct {
	Age       int    `njson:"age"`
	Name      string `njson:"name"`
	Address   string `njson:address"`
}
type re_student struct {
       Per_1    re_person `njson:"re_person"`
	   Year_of_study  int  `njson:"year_of_study"`
}
type person struct {
	Age       int    
	Name      string 
	Address   string 
}
type student struct {
       Per_1    person 
	   Year_of_study  int 
}
func func_mar()  string{
     var m student
	 m.Per_1.Age = 20
	 m.Per_1.Name = "gopi"
	 m.Per_1.Address = "Nrt"
	 m.Year_of_study = 4
      
	 
    jsonString, err := json.Marshal(m)
	if err != nil{
	   fmt.Println(err)
	 }
	 //fmt.Printf("%T\n", jsonString)
	 fmt.Println("marshal", string(jsonString))
	encodedStr := base64.StdEncoding.EncodeToString([]byte(jsonString))
    fmt.Println("Encoded:", encodedStr)
	fmt.Printf("%T\n", encodedStr)
	   return encodedStr
}
func func_unmar(encodedStr string)  string{

	// Decoding may return an error, in case the input is not well formed
    decodedStr, err := base64.StdEncoding.DecodeString(encodedStr)
    if err != nil {
        panic("malformed input")
    }
	
	 // convert json to struct
    var x =string(decodedStr)
	fmt.Println("Decoded & converted String", x)
	stBytes := []byte(x)
    var s re_student
    json.Unmarshal(stBytes, &s)
    fmt.Println("Unmarhal ", s)
	return s
}

func main() {
    var en1 = func_mar()
	func_unmar(en1)
	
   /*   var jsonString = func_mar()
    
	*/  
}