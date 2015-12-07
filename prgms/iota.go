package main
import "fmt"
func main(){
var a int =15
var b=20
var c= a|b
var d= a&b
fmt.Print(a)
fmt.Printf("\t%b\n",a)
fmt.Print(b)
fmt.Printf("\t%b\n",b)
fmt.Print("after bitwise or operation\n")
fmt.Print(c)
fmt.Printf("\t%b\n",c)
fmt.Print("after bitwise and operation\n")
fmt.Print(d)
fmt.Printf("\t%b",d)
fmt.Println("\nafter not operation")
fmt.Print(^a)
fmt.Printf("\t%b",a)
}
