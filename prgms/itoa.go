package main
import(
 "fmt"
)
func main(){
var i int
i = 22
fmt.Print(i)
fmt.Printf("\n%T",i)
var f float64
f = float64(i)+0.523
fmt.Println("\n",f)
fmt.Printf("%T",f)
}
