package main
import (
"fmt"
"reflect"
)
func main(){
var s string = "kala"
var i int = 56
var f float64 =32.234
fmt.Println(reflect.TypeOf(s))
fmt.Println(reflect.TypeOf(i))
fmt.Println(reflect.TypeOf(f))
}