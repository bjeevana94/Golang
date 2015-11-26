package main
import "fmt"

type cstmr struct{
name string
id   string
sal  float64
age  int
}
func main(){
c1 := new(cstmr)
c1.name = "kala"
c1.id = "12d34"
c1.sal = 10000.2
c1.age = 24
fmt.Println(c1.name)
fmt.Println("before updating", *c1)
c1.name = "ashu"
fmt.Println("after updating a value")
fmt.Println(c1.name)
fmt.Println(*c1)
}
