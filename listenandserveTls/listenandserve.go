package main
import(
	"fmt"
	"net/http"
	"log"
	"html/template"
	"github.com/kabukky/httpscerts"
	 "github.com/gorilla/mux"
)
func init(){


	tpl = template.Must(template.New("").ParseGlob("templates/*.tmpl"))
	
}
func main(){
    err := httpscerts.Check("cert.pem", "key.pem")
	if err != nil {
        err = httpscerts.Generate("cert.pem", "key.pem", "gitty.us")
        if err != nil {
            log.Fatal("Error: Couldn't create https certs.")
        }
    }
		myRouter := mux.NewRouter().StrictSlash(true)
		myRouter.HandleFunc("/",Home)
		err = http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", myRouter)
    	 log.Fatal(err)
}
func Home(w http.ResponseWriter, req *http.Request ){
	tpl.ExecuteTemplate(w, "index.tmpl", nil)
}