package main
import(
	"fmt"
	"net/http"
	"log"
	"html/template"

	 "github.com/gorilla/mux"
)
func init(){


	tpl = template.Must(template.New("").ParseGlob("templates/*.tmpl"))
	
}
func main(){

		myRouter := mux.NewRouter().StrictSlash(true)
		myRouter.HandleFunc("/",Home)
		http.ListenAndServe(":8080",myRouter)
}
func Home(w http.ResponseWriter, req *http.Request ){
	tpl.ExecuteTemplate(w, "index.tmpl", nil)
}