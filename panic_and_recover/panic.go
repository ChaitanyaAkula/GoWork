package main
import (
	"fmt"

)
func main(){
	//var action int
	division()
}
func division(){
	var a,b int
	defer func(){
		if action:=recover();action!=nil{
			fmt.Println("recovered",action)
		}
		
	}()
	fmt.Print("enter a ,b values :")
	fmt.Scanln(&a,&b)
	g(a,b)

}
func g(x,y int){
 	var	value int=0
	if y!=0 {
		
		value= x/y

	}else if y==0 {
		fmt.Println("panicking")
		panic(fmt.Sprintf("y value is 0"))

	}
	fmt.Println(value)
	

}