package main

import (
	"fmt"
	"gomodule/src/service"
)

func main() {
	stu:=service.Student{"chen",11}
	stu.Name="41"
	stu2:=service.Student{}
	stu2.Age=13
	stu2.Name="chen"
	fmt.Println(stu2)
	fmt.Println(stu)
}
