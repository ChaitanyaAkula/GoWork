package main
import(
	"fmt"
	"net/http"
	"log"
	"text/template"
	"bufio"
	"strings"
	"net"
)
type webpage int
var tpl *template.Template
func(wp webpage)ServeHTTP(w http.ResponseWriter, req *http.Request){
	//fmt.Fprintln(w,"My first WebPage")
	err:=req.ParseForm()
	if err!=nil{
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w,"index.html",req.Form)
}
func init(){
	tpl =template.Must(template.ParseFiles("index.html","login.html"))
}
func main(){
	var n webpage
	http.ListenAndServe(":8080",n)
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
	
		go handle(conn)



	}
}
func handle(conn net.Conn){
	
	defer conn.Close()
	request(conn)

}
func request(conn net.Conn){
	i:=0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		line:= scanner.Text()
	//	fmt.Println(line)
		if i==0{
			mux(conn,line)

		}
		if line==""{
			break //headers are done
		}
		i++
	}
}
func mux(conn net.Conn,ln string){
	m:=strings.Fields(ln)[0]    //method
	u:=strings.Fields(ln)[1]    //URL

	//multiplexer

	if m=="GET" && u=="/"{
		index(conn)
	}
	if m=="GET" && u=="/login"{
		login(conn)
	}

}
func index(conn net.Conn){
 body:=`<!DOCTYPE html>
 <html>
	 <head><title>Welcome</title></head>
	 <body>
		 <p>Home</p>
		 <a href="/login">Login</a><br>
	 </body>
 </html>`
 fmt.Fprintln(conn,body)
}
func login(conn net.Conn){
	body:=`<!DOCTYPE html>
 <html>
	 <head><title>Welcome</title></head>
	 <body>
		 <p>Login</p>
		 <a href="/index">Home</a><br>
	 </body>
 </html>`
 fmt.Fprintln(conn,body)
}