package main
import(
	//"fmt"
	//"html/template"
	"text/template"
	"log"
	"os"
	"strings"
)
var tpl *template.Template
var fm = template.FuncMap{
	"uc":strings.ToUpper,
	"ft":firstThree,
}
func init(){
	//tpl=template.Must(template.ParseGlob("gohtml/*"))
	tpl=template.Must(template.New("").Funcs(fm).ParseGlob("gohtml/*"))
}
func firstThree(s string)string{
	s= strings.TrimSpace(s)
	s=s[:3]
	return s


}
type currency struct{
	 
	Country string
	Currency string

}
type president struct{
	Name string

}
type country struct{
	Capital []currency
	Pres    []president
}

func main(){
	
	mapv := map[string]string{
		"india" : "delhi",
		"South korea ": "seoul",
		"USA"         :"Washington,D.C",

	}
	sindia := currency {
		Country : "India",
		Currency : "",
	}
	skorea := currency{
		Country: "Korea ",
		Currency: "Won",
	}
	susa:=currency{
		Country :"USA",
		Currency: "Dollar",
	}
	pindia := president{
		Name: "Ram Nath Kovind",
	}
	pkorea := president{
		Name: "Moon Jae-in",
	}
	pusa:=president{
		Name: "Donald trumph",
	}

	currencies := []currency{sindia,skorea,susa}
	presidents := []president{pindia,pkorea,pusa}
	data := country{
		Capital : currencies ,
		Pres     : presidents,
	}

	err := tpl.ExecuteTemplate(os.Stdout,"map.gohtml",mapv)
	//err = tpl.ExecuteTemplate(os.Stdout,"struct.gohtml",sindia)
	//err = tpl.ExecuteTemplate(os.Stdout,"struct.gohtml",skorea)
	err = tpl.ExecuteTemplate(os.Stdout,"struct.gohtml",data)
	if err!=nil{
		log.Fatalln(err)
	}

}