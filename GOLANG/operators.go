package main
import (
	"fmt"
)
func main(){
	
	fmt.Println("")
	var a,b int
	fmt.Println( "Enter a and b vlaues =")
	fmt.Scanln(&a,&b)
	fmt.Println("")
	fmt.Println( "a value =",a," ","b value=",b)
	fmt.Println("")
	fmt.Println( "Enter firstname and lastname =")
	var fname,lname string
	fmt.Scanln(&fname,&lname)
	fmt.Println("")
	fmt.Println("firstname=",fname ,"lastname=",lname)
	fmt.Println("")

	fmt.Println("Arithmetic operators +,-,*,/,%,concatination")
	fmt.Println("")

	defer comparison(a,b)
	fmt.Println("")

	fmt.Println(arithmetic(a,b))
	fmt.Println(fname+lname)
	fmt.Println("")

	fmt.Println("Bitwise operators &, |, ^, &^ ")
	fmt.Println(bitwise(a,b))
	fmt.Println("")

	fmt.Println("Logical operators")
	fmt.Println(logicalop(a,b))
	fmt.Println("")

	fmt.Println("Assignment operators")
	fmt.Println(assignment(a,b))
	fmt.Println("")
	
	fmt.Println("Comparison operators")
	

	
}
func arithmetic(x int,y int)(int,int,int,int,int){
  var sum,diff,mul,div,remainder int=0,0,0,0,0
  sum=x+y
  diff=x-y
  mul=x*y
  div=x/y
  remainder=x%y
 return sum,diff,mul,div,remainder
}
func bitwise(x int,y int)(int,int,int,int){
	var z1,z2,z3,z4 int
	z1= x&y
	z2 = x|y
	z3 = x^y
	z4 = x&^y
	return z1,z2,z3,z4
}
func logicalop(n1 int,n2 int)int{

	var num1 int =n1
	var num2 int =n2
	if(num1!=num2 && num1<=num2){
		fmt.Println(true)
	 }
	 if(num1!=num2 || num1<=num2){
			fmt.Println(true)
	 }
	 if(!(num1==num2)){
		fmt.Println(true)

	}
	return 0
}
func assignment(n1 int,n2 int)int{
	var x int=n1
	var y int=n2
	x+=y
	fmt.Println("add AND assignment operator(x+=y) : ",x)
	fmt.Println("")
	x-=y
	fmt.Println("subtract AND assignment operator(x-=y) : ",x)
	fmt.Println("")
	x*=y
	fmt.Println("Multiply AND assignment operator(x*=y) : ",x)
	fmt.Println("")
	x/=y
	fmt.Println("Divide AND assignment operator(x/=y) : ",x)
	fmt.Println("")
	x%=y
	fmt.Println("modulus AND assignment operator(x%=y) : ",x)
	fmt.Println("")
	return 0
}
func comparison(x int,y int) {
	fmt.Println("Used defer statement to call Comparison function ")
	fmt.Println("")
	fmt.Println(x==y)
	fmt.Println(x!=y)
	fmt.Println(x>y)
	fmt.Println(x<y)
	
	
}