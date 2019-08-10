package main
import(
	"fmt"
	"unsafe"
)
func main(){
   a:=struct{
	   A float32 
	   B string

   }(0,"go")
   fmt.Println(unsafe.Sizeof(s))
}