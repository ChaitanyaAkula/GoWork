package main
import(
	"fmt"
)
func main(){
	var student = map[string]int{}
	var student1 = map[string]int{
		"subj1" : 80,
		"subj2" : 90,
		"sub3" : 99,
	}
	var capital = make(map[string]string)
	capital["AP"] = "Amaravathi"
	capital["Telengana"]="Hyderabad"
	capital["TN"]= "Chennai"

	fmt.Println(student)
	fmt.Println(student1)
	fmt.Println(capital)
	fmt.Println("")
	delete(capital,"TN")
	fmt.Println(capital)
	student["chaitu"]=10
	fmt.Println("")
	fmt.Println(student)
	for key,value:=range student1{
		fmt.Println(key,"-->",value)
	}


}

