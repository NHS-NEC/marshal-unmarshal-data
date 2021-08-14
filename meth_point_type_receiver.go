// Go program to illustrate pointer receiver

package main

import (
	"fmt"
)
type student struct {
	stu_no int
	stu_name string
	stu_marks int
}
func (st *student) fun_mod(stu_name string) {
	(*st).stu_name = stu_name
}
func (st student) fun_show() {
	fmt.Println("Student Number: ", st.stu_no)
    fmt.Println("Student Name: ", st.stu_name)
    fmt.Println("Student Marks: ", st.stu_marks) 
}
func main() {
    
	student1 := student {
	stu_no    : 519,
	stu_name  : "NEC",
	stu_marks : 500,
	}
	student1.fun_show()
	// Creating a pointer
    p := &student1
  
    // Calling the show method
    p.fun_mod("STR")
	student1.fun_show()
}
/*  Output
 go run f1_swap.go
Student Number:  519
Student Name:  NEC
Student Marks:  500
Student Number:  519
Student Name:  STR
Student Marks:  500
*/
