package main
import "fmt"
func main(){
var rad float64
var pm float64
fmt.Println("enter radius value:")
fmt.Scanln(&rad)
pm = 2*3.14*rad
fmt.Println("perimeter of circle:", pm)
}