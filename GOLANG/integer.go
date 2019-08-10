package main
import(
	"fmt"
)
func main(){
	const name string="chaitanya"	
	//unsigned int
	var a uint8
	a=200
	var b uint16= 25560
	var c uint32= 1258905
	var d uint64= 56897885589
	fmt.Println("unsigned integer values")
	fmt.Println(" ")
	fmt.Println(a,b,c,d)

	//signed integers
	var n1 int8
	n1=-20
	var n2 int16= -25560
	var n3 int32= -1258905
	var n4 int64= -56897885589
	fmt.Println(" ")

	fmt.Println("signed integer values")
	fmt.Println(" ")
	fmt.Println("n1=",n1,"n2=",n2,"n3=",n3,"n4=",n4)
	fmt.Println(" ")

	fmt.Println("n1*-1 =",n1*-1,"n2*-1= ",n2*-1,"n3*-1= ",n3*-1,"n4*-1= ",n4*-1)

	//Float type

	var f1 float32 = 564154.255
	var f2 float64 = 85566.256
	var f3 complex64 = 4654654.256
	var f4 complex128 = 45464564164.2568
	
	fmt.Println(" ")
	fmt.Println("Floating-point values and f3 and f4 are complex numbers")
	
	fmt.Println(" ")
	fmt.Println(f1,f2,f3,f4)
	
	//Boolean type
	var b1,b2 bool
	b1=true
	b2=false
	fmt.Println("Boolean type")
	fmt.Println(" ")
	fmt.Println(b1,b2)
	fmt.Println(" ")
	//fmt.Println("String")
	var city string = "Hyderabad"
	var country="INDIA"
	fmt.Println(city,country)
	
	fmt.Println("Hello,World! this is ",name)
}