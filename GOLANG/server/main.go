package main
import(
	"fmt"
	"io"
	"log"
	"net"
	"bufio"
)
func main(){
	li,err := net.Listen("tcp",":8080")
	if err!=nil{
		log.Panic(err)
	}
	defer li.Close()
	for{
		conn,err:=li.Accept()
		if err!=nil{
			log.Panic(err)
		}
		io.WriteString(conn,"\n Hello,World! This is chaitanya\n")
		fmt.Fprintln(conn,"Creating a server  ")
		fmt.Fprintf(conn,"%v","Good afternoon ")
		go handle(conn)



	}
}
func handle(conn net.Conn){
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		line:= scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn,"I heard you say : %s\n",line)
	}
	defer conn.Close()
}
