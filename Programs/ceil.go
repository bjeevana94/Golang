package main
import(
"fmt"
"math"
)
func main(){
var i float64
fmt.Println("enter floating point number:")
fmt.Scan(&i)
fmt.Println("least integer of your float value", math.Ceil(i))
}