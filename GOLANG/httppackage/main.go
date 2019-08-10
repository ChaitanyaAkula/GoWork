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
	//var n webpage
	http.ListenAndServe(":8080",nil)
	

	}
}
