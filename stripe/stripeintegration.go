package main
import(
	"fmt"
	"net/http"
	"log"
	"github.com/ChaitanyaAkula/GoWork/database"
    "github.com/ChaitanyaAkula/GoWork/cookies"	

)
func stripePayment1(w http.ResponseWriter,req *http.Request){
	var balance int
	db:=dbconnection.Connection()
	defer db.Close()
	comid,_:=req.Cookie("Companyid")
	 rows,_:=db.Query("select balance from companylogin WHERE idcompanylogin=?",comid.Value)
		 for rows.Next(){
			 err:=rows.Scan(&Companybalance)
				 
			 if err!=nil{
				 log.Fatalln(err)
			 }
		 }	
Stripetoken=req.FormValue("stripeToken")
	cookies.SetCookie(w,"stripeToken",Stripetoken)
	

	
insForm,err1 := db.Prepare("UPDATE companylogin SET balance= ?,stripeToken=? WHERE idcompanylogin=? ")
if err1!=nil{
			log.Fatalln(err1)
		 }
		 _,err= insForm.Exec(balance,Stripetoken,comid.Value)
		if err!=nil{
						 log.Fatalln(err)
			 }

	 stripe.Key = STRIPE_SECRETKEY

	 params := &stripe.CustomerParams{
	      Email: stripe.String("chaitanyaakula1910@gmail.com"),
     }
	 params.SetSource(Stripetoken)
	cus, _ := customer.New(params)	
	fmt.Println(cus)	
	http.Redirect(w,req,"/postconfirmation",307)



}