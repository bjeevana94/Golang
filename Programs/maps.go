package main
import "fmt"
func main(){
map1 := map[int]string{
0: "my world",
1: "tanu",
2: "ashu",
3: "tanusha",
}
fmt.Println(len(map1))
for _, val :=range map1{
fmt.Println(val)
}
delete(map1,3)
fmt.Println("after deleting")
fmt.Println(map1)
fmt.Println("\nchecking for 3rd element in map")
if _, exists := map1[3]; exists{
fmt.Println("value exists")
fmt.Println(map1)
}else{
fmt.Println("value doesnot exists:")
fmt.Println(map1)
}
}