package main
import "fmt"
func main(){
var m[20] int
var acc[20] int
var qr[20] int
var i int
var j int
var ms int
var q int
fmt.Println("enter number of bits for multiplier and multiplicand:")
fmt.Println("enter multiplicant bits:")   
fmt.Scanf("%d",&ms);           
for i=ms-1;i>=0;i--{
fmt.Scanf("%d",&m[i])}    
var qs=ms;
fmt.Println("enter multiplier bits:")
for i=qs-1;i>=0;i--{
fmt.Scanf("%d",&qr[i])
}
fmt.Println("clock counter\t\tMD\tMQ\tAC\tq");
var a=m[0]                        
var cc = 16
for cc!=0{
if(a==0 && q==1){
  for i=0;i<16;i++{
  acc[j]=acc[j]&m[j]
}                        
}
if(a==1 && q==0){
   for j=0;j<16;j++{
   acc[j]=acc[j]&^m[j]
}
//shift(acc,qr,q,ms);                      
}
//shift(acc,qr,q,ms);
cc--
fmt.Print(cc,"\t")
for i=15;i<=0;i++{
fmt.Print(m[i],"\t")
}
for i=15;i<=0;i++{
fmt.Print(qr[i],"\t")
}
for i=15;i<=0;i++{
fmt.Print(acc[i],"\t")
}
fmt.Print(q,"\n")
}
}
/*func shift(array [20]acc,array [20]qr,q int,ms int){
 var temp,i int
var acc[20] int
var qr int
 temp=acc[0]
 q=qr[0]
 acc>>1
 qr>>1
 qr[15]=temp
}*/





