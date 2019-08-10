package main
import(
	"fmt"
)
func main(){
	var a,b,c,d,e int
	fmt.Printf("enter 5 values :")
	fmt.Scanln(&a,&b,&c,&d,&e)

		variadic("Hyderabd","vijayawada","chenai")
		variadic("chaitanya")
	   calculation("rectangle",a,b)	 
	   sum(a,b,c,d,e)
}
func variadic(s ...string){
	fmt.Println(s)
	//fmt.Println(s[1])
	//fmt.Println(s[2])
}
func calculation(str string,x ...int){
	 area:=1					
	
		for _,v:= range x{
			if str=="rectangle"{
				area *=v
			}
			
		}
		
			fmt.Println("area is \t : ",area)
		}

func sum(values ...int){
	result:=0
	for _,v:= range values{
		result+=v;
		
	}
	fmt.Println("total:",result)
}
	      
	
		
