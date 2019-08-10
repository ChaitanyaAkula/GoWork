package main
import(
	"fmt"
	"time"
	"math/rand"
	"sync"
	"runtime"
)
var wg sync.WaitGroup
var mutex sync.Mutex
var counter int
type testConcurrency struct{
	min int
	max int
	country string
} 
func printConcurrency(test *testConcurrency){
    for i:=test.max;i>test.min;i--{
		time.Sleep(3*time.Millisecond)
		fmt.Println(test.country)
	} 
    
    wg.Done()
}
func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}



func f(n int){
	for i:=0;i<10;i++{
		fmt.Println("foo",":",i)
		amt:=time.Duration(rand.Intn(10))
		time.Sleep(time.Millisecond+amt)
	    wg.Done()
	}
}
func g(n int){
	for i:=0;i<10;i++{
		fmt.Print("goo",":",i)
		//amt:=time.Duration(rand.Intn(250))
		time.Sleep(time.Duration(3*time.Millisecond))
		//mutex.Lock()
		counter++
		fmt.Println("counter : " ,counter)
		//mutex.Unlock()
		wg.Done()
	}
}
func main(){
	india:= new(testConcurrency)
	japan:= new(testConcurrency)
	china:=new(testConcurrency)
    //var china = testConcurrency{10,15,"China"}
	japan.min=5
	japan.max=6
	japan.country="Japan"
	india.min=50
	india.max=50
	india.country="India"
	china.min=2
	china.max=3
	china.country="china"
	go printConcurrency(india)
	go printConcurrency(japan)
	go printConcurrency(china)
	wg.Add(3)
	wg.Wait()

	wg.Add(2)
	go f(0)
	for i:=0;i<10;i++{
		go g(i)
	}
	wg.Wait()
	//var japan testConcurrency

	//var input string
	
  // fmt.Scanln(&input)
}