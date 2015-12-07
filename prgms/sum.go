package main
import "fmt"
func main(){
var i int
var sum=0
for i=1;i<100;i++{
if i%3==0 || i%5==0{
sum=sum+i
}
}
fmt.Println("sum of all multiples of 3 or 5:",sum)
}