package main

import (
	"fmt"
)

func fun_add(x,y int) int {
   return x + y
}

func main() {
   var a=70
   var b=30

   fmt.Println(fun_add(a,b))

}