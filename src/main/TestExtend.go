package main
import "fmt"

type BaseNum struct{
	num1 int
	num2 int
}
type Add struct {
	BaseNum
}
type Sub struct {
	BaseNum
}
func(a *Add)opt()(value int){
	return a.num1+a.num2
}
func(a *Sub)opt()(value int){
	return a.num1+a.num2
}
type Opter interface {
	opt() int
}
func MultiState(o *Opter)(value int) { //多态定义, 可以简单理解为以接口作为形参的函数, 方便学习

	return value
}

func main() {

// 继承
data := BaseNum{2,3}
var a =Add{data}
var b =Sub{data}
var i Opter
i=&a
value:=i.opt()
fmt.Println(value)
i=&b
value1:=i.opt()
fmt.Println(value1)



}