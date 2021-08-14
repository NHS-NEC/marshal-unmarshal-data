package main
  
import "fmt"
  
// Student structure
type student struct {
	stu_no int
	stu_name string
	stu_marks int
}

  
// Method with a pointer
// receiver of author type
func (a *student) show_1(stu_no int) {
    (*a).stu_no = stu_no
}
  
// Method with a value
// receiver of author type
func (a student) show_2() {
  
    a.stu_name = "Gourav"
    fmt.Println("Author's name(Before) : ", a.stu_name)
}
  
// Main function
func main() {
  
    // Initializing the values
    // of the author structure
    res := student{
        stu_no: 501,
	    stu_name: "str",
	    stu_marks:400,
    }
  
    fmt.Println("Student Noe(Before): ", res.stu_no)
  
    // Calling the show_1 method
    // (pointer method) with value
    res.show_1(510)
    fmt.Println("Branch Name(After): ", res.stu_no)
  
    // Calling the show_2 method
    // (value method) with a pointer
    (&res).show_2()
    fmt.Println("Author's name(After): ", res.stu_name)
}
/* OUTPUT
go run f1_swap.go
Student Noe(Before):  501
Branch Name(After):  510
Author's name(Before) :  Gourav
Author's name(After):  str
*/
