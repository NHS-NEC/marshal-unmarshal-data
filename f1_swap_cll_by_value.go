package main

import (
	"fmt"
)

func fun_swap(x,y int) {
   var temp int

   temp = x
   x=y
   y=temp
}

func main() {
   var a=70
   var b=30

   fmt.Println("Before Swaping", a, b)
   fun_swap(a,b)
   fmt.Println("Aftere Swaping", a, b)

}
/*
go run f1.go
Before Swaping 70 30
Aftere Swaping 70 30
*/