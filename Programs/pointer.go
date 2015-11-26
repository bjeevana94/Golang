package main
import "fmt"
func main(){
var i int=20
var a *int =&i
fmt.Println("value if i:",i)
fmt.Println("address of i:",&i)
fmt.Println("pointer value :",*a)
}