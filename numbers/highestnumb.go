package main
import(
	"fmt"
	"sort"
	"sync"
	"time"
)
var wg sync.WaitGroup
func main(){
//	a := []int{6,8,9,12,56,1}
	arr := make([]int,5,20)
	wg.Add(5)
	fmt.Println("enter the elements : ")
	for i:=0;i<len(arr);i++{
		fmt.Printf("elements a[%d] :",i)
		fmt.Scanln(&arr[i])
	}
	fmt.Println("Unsorted array",arr)
	
	sort.Ints(arr)
	fmt.Println("sorted array:",arr)
	go thirdHighest(arr)
	func (){
		for i:=0;i<10;i++{
			fmt.Println("foo:",i)
		}
	
	}()
	go secondHighest(arr)
		
	wg.Wait()
}
func secondHighest(arr []int){
	//fmt.Println("testing")
	time.Sleep(time.Duration(3*time.Millisecond))
	fmt.Println("Second highest number",arr[len(arr)-2])
	
	wg.Done()


}
func thirdHighest(arr []int){
	
	time.Sleep(time.Duration(5*time.Millisecond))
	fmt.Println("Third highest number",arr[len(arr)-3])
	
	wg.Done()


}
