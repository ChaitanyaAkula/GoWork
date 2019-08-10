package main
import(
	"fmt"
)
func main(){
	var n,a,b int
	fmt.Println("enter a number :")
	fmt.Scanln(&n)
	primenumber(n)
	fmt.Println("enter 2 numbers :")
	fmt.Scanln(&a,&b)
	primeupto(a,b)

	
}
func primenumber(n int){
	flag :=true
	for i:=2;i<n/2;i++{
		if n%i==0{
			flag= false
			break
		}
	}
	if flag==false{
		fmt.Printf("n=%d is not a prime number",n)
		fmt.Println(" ")
	}else{
		fmt.Printf("n=%d is a prime number",n)
		fmt.Println("")
	}

}
func primeupto(a,b int){
	flag :=true
	for i:=a;i<=b;i++{
		if a%i==0{
			flag= false
			break
		}
	}
	if flag==true{
		for i:=a;i<=b;i++{
			fmt.Println("prime number from %d to %d is :%d ",a,b,i)
		}

	}
	
}