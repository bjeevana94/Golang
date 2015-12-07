package main
import "fmt"
func main(){
var a,b int
fmt.Println("enter values of a,b:")
fmt.Scanf("%d%d",&a,&b)
var c int =a%b
fmt.Println("after division:",c)
}