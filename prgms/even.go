package main
import "fmt"
func main(){
var i int
fmt.Println("even numbers from 1 to 1000")
for i=1;i<=1000;i++{
if(i%2==0){
fmt.Print(i)
}
fmt.Print(" ")
}
}