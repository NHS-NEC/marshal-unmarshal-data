/* Encoding Decoding with mongoDB  */
package main

import (
	"encoding/json"
	"encoding/base64"
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson"
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
func func_unmar(encodedStr string, client *mongo.Client)  re_student{

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
	 collection := client.Database("test").Collection("trainers")
    ash :=  s
    insertResult, err := collection.InsertOne(context.TODO(), ash)
    if err != nil {
    log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	 
	return s
}
func func_mong_connect() *mongo.Client{
      // Set client options
      clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

      // Connect to MongoDB
      client, err := mongo.Connect(context.TODO(), clientOptions)

      if err != nil {
             log.Fatal(err)
        }
       // Check the connection
       err = client.Ping(context.TODO(), nil)

       if err != nil {
             log.Fatal(err)
          }
       fmt.Println("Connected to MongoDB!")
	   //fmt.Print("Type client %T", client)
	   return client
}
func main() {
    var en1 = func_mar()
	client1 := func_mong_connect()
	
	fmt.Println("In The MAIN", func_unmar(en1, client1))
	
}