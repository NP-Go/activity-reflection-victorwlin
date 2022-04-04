package main

import (
	"fmt"
	"reflect"
)

//Declare Inspection Function
func inspector(x interface{}) {
	typ := reflect.TypeOf(x)
	fmt.Println("")
	fmt.Println(x)
	fmt.Println(typ.Name())
	fmt.Println(typ.Kind())
}

func main() {
	//Insert code here
	str := "This is a string"
	integer := 12345
	float := 1.2345
	boolean := true

	inspector(str)
	inspector(integer)
	inspector(float)
	inspector(boolean)
}
