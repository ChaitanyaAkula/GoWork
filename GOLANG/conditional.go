package main
import(
	"fmt"
     "time"
)
func main(){
	today := time.Now()
	var t int=today.Day()
	/*today1 := time.Date(2019,4,2,11,15,0,0,time.UTC)
	
	year,month,day :=today1.Date()
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)*/
	fmt.Println("current date and time is :",today.String())
	fmt.Println(" ")
	fmt.Println(today.Format("01/04/2019 Monday "))
	fmt.Println(" ")


	//var t int =1
	switch  {
	//case 1:
         //fmt.Println("date is 1")
		//fallthrough
	case t<10:
		fmt.Println("today date is",t)
		fallthrough
	case t==10:
		
		fmt.Println("Month is",today.Month())
		fmt.Println("Weekday is",today.Weekday())
		fmt.Println("year is",today.Year())
	
	default:
		fmt.Println("NO information")
		
	}
var rows int
fmt.Println("Enter number of rows :")
fmt.Scanln(&rows)
fmt.Print("\n")
fmt.Println("format1")

format1(rows)
fmt.Print("\n")
fmt.Println("format2")
format2(rows)
}
func format1(n int){
	for i:=0;i<n;i++{
		for j:=0;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
func format2(n int){
	for i:=n;i>0;i--{
		for j:=0;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}