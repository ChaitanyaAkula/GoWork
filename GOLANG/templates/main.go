package main
import(
	"fmt"
	//"html/template"
	"text/template"
	"log"
	"os"
)
func main(){
	message:= "Hello, World !"
	html:=`
	<!DOCTYPE html>
	<htmllang="en">
		 <head>
		 	<meta charset="UTF-8">
			 <title>My First WebPage</title></head>
			 <body>
			 <h1> `+ message+ `<h1>
			 </body>

	</html>
	`
	fmt.Println(html)
	s:=[]string{"chaitu","krish","akula"}


	tpl,err:= template.ParseGlob("gohtml/*")
	if err!=nil{
		log.Fatalln(err)
	}

//	tpl,err = tpl.ParseFiles("chaitu.xht")
//	if err!=nil{
	//	log.Fatalln(err)
	//}
	//err=tpl.Execute(os.stdout)
	nf,err := os.Create("index2.html")
	if err!=nil{
		log.Fatalln(err)
	}
	defer nf.Close()
//n:="chaitu"
	err=tpl.ExecuteTemplate(os.Stdout,"chaitu.gohtml",`chaitu`)
	err=tpl.ExecuteTemplate(os.Stdout,"index.gohtml",`this is my First WebPage`)
	err=tpl.ExecuteTemplate(os.Stdout,"table.gohtml",s)
	if err!=nil{
		log.Fatalln(err)
	}
	err=tpl.Execute(nf,nil)
	//if err!=nil{
	//	log.Fatalln(err)
	//}

}