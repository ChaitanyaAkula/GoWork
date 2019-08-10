package main
import(
	//"fmt"
	"net/http"
	"log"
	"text/template"
	//"bufio"
	//"strings"
	//"net"
)

var tpl *template.Template

	

type person struct{
	Fname string
	Surname string
	Emailid string
	Phonenumber string
}
func init(){
	tpl =template.Must(template.ParseFiles("index.html","home.html","login.html"))
}
func main(){
	//var n webpage
	http.HandleFunc("/",index)
	http.HandleFunc("/login",login)

	http.ListenAndServe(":8080",nil)
}
func index(w http.ResponseWriter, req *http.Request ){

			err :=tpl.ExecuteTemplate(w,"home.html",nil)
			if err!=nil{
				log.Fatalln(err)
			}
			
	}

func login(w http.ResponseWriter, req *http.Request ){
	       fn:=req.FormValue("first")
			sn:=req.FormValue("surname")
			id:=req.FormValue("email")
			numb:=req.FormValue("phonenumber")
			p:=new(person)
			p.Fname =fn
			p.Surname=sn
			p.Emailid=id
			p.Phonenumber=numb
			err :=tpl.ExecuteTemplate(w,"login.html",&p)
			if err!=nil{
				log.Fatalln(err)
			}
}