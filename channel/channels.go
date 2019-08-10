package main
import(
	"fmt"
	"sync"
	"time"
	"math/rand"
)
var wg = sync.WaitGroup{}
func main(){
	
	ch:= make(chan int,50)
	projects := make(chan string,10)
	wg.Add(10)
	go func(ch<- chan int){  // recieve-only channel
	
		for val:= range ch{
			fmt.Println("value :",val)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int){ //send-only channel
		i:=10
		for ;i<=15;i++{
			ch<-i
		
			
		}
		
		close(ch)
		wg.Done()

	}(ch)
	//wg.Wait()

	for i:=0;i<=5;i++{
		go employee(projects,i)
	}
	for j :=1; j <= 10; j++ {
        projects <- fmt.Sprintf("Project :%d", j)
	}
	close(projects)
	wg.Wait()
}

func employee(projects chan string,employee int){
for{
		project,result :=<-projects

	if result==false{
		fmt.Printf("Employee :%d\n",employee)
		return
	}
	fmt.Printf("Employee : %d : Started   %s\n", employee, project)
	sleep := rand.Int63n(50)
	time.Sleep(3 * time.Millisecond)
	fmt.Println("\nTime to sleep",sleep,"ms\n")
    fmt.Printf("Employee : %d : Completed %s\n", employee, project)

	wg.Done()
}
}
