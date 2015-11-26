package main
import "fmt"
func main(){
var a=10
var b=12
var c=a|b
var d=a&b
var e=a&^b
fmt.Printf("%b\n",a)
fmt.Printf("%b\n",b)
fmt.Printf("after bitwise OR operation:%b\n",c)
fmt.Printf("after bitwise AND operation:%b\n",d)
fmt.Printf("after NOT operation:%b",e)
fmt.Printf("after shift operation:%b",b>>1)
//fmt.Printf("after shift operation:%b",b[0])
}