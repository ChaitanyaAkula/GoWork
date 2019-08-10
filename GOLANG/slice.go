package main
import(
	"fmt"
)
func main(){
	var a = make([]int,10)
	var b = make([]int,15,20)
	 c := new([25]int)[0:15]
	var stringslice = []string{"chaitanya","akula"} 
	fmt.Printf("slice a :\tlength: %v \t capacity : %v\n",len(a),cap(a))
	fmt.Printf("slice b :\tlength: %v \t capacity : %v\n",len(b),cap(b))
	fmt.Printf("slice c :\tlength: %v \t capacity : %v\n",len(c),cap(c))
	fmt.Println(stringslice)
	a= append(a,10,20,30,40,50)
	stringslice=append(stringslice,"prathyu","vineel","phani")
	fmt.Println(a)
	fmt.Println(stringslice)
	var str1= make([]string,10,20)
	copy(str1,stringslice)
	fmt.Println(str1)
	fmt.Printf("after trim :%s",stringslice[:3])
	str1=append(str1,stringslice[:1]...)
	fmt.Println(str1)
	fmt.Println(stringslice[len(stringslice)-2])


}