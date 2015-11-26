package main
import "fmt"
func main(){
var f float64
f = 32.45
fmt.Println("before conversion:",f)
fmt.Println("after conversion",int(f))
}