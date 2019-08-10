package main
import(
	"fmt"
)
type student = struct{
	Name  string
	Age   int
	Grade string
}
type shape struct{
	length int
	breadth int
	geometry struct{
		area int
		perimeter int
	}
}
func main(){
	var rectangle shape
	
	rectangle.length= 20
	rectangle.breadth=30
	rectangle.geometry.area =rectangle.length*rectangle.breadth
	rectangle.geometry.perimeter =2*(rectangle.length+rectangle.breadth)
	

	fmt.Println(student{"chaitu",23,"A"})
	fmt.Println("Area of the rectangle is :",rectangle.geometry.area)
	fmt.Println("Perimeter of the rectangle :",rectangle.geometry.perimeter)

	rect1:= &shape{length:10}
	rect1.geometry.area=rect1.length*rect1.breadth
	fmt.Println("area of react1: ",rect1.geometry.area)

	var rect2 = shape{length:30,breadth:30}
	rect2.geometry.area =rect2.length*rect2.breadth
	fmt.Println("area of rect2 :",rect2.geometry.area)

	var rect3 =new(shape)
	rect3.length=5
	rect3.breadth=6
	rect3.geometry.area=rect3.length*rect3.breadth
	fmt.Println("area of rect3 :",rect3.geometry.area)
}