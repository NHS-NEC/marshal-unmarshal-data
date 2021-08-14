package main
import (
    "encoding/json"
    "fmt"
//    "os"
)
//We’ll use these two structs to demonstrate encoding and decoding of custom types below.
type response1 struct {
    Page   int
    Fruits []string
}
//Only exported fields will be encoded/decoded in JSON. Fields must start with capital letters to be exported.
type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}
func main() {
//frst we’ll look at encoding basic data types to JSON strings. Here are some examples for atomic values.
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))
	
	intB, _ := json.Marshal(1)
    fmt.Println(string(intB))
	
	fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))
	
	strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))
	
	//And here are some for slices and maps, which encode to JSON arrays and objects as you’d expect.
	slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))
    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))
	
	//The JSON package can automatically encode your custom data types. 
	//It will only include exported fields in the encoded output and will 
	//by default use those names as the JSON keys.
	res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))
	
	//You can use tags on struct field declarations 
	//to customize the encoded JSON key names. Check the 
	//definition of response2 above to see an example of such tags
	res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))
	
	//Now let’s look at decoding JSON data into Go values. 
	//Here’s an example for a generic data structure.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
}