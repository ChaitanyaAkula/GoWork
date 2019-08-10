package main
import(
	"fmt"
	"net/http"
	"log"
	"html/template"
	//"bufio"
	//"net/smtp"
	"math/rand"
	"encoding/json"
     "github.com/gorilla/mux"
	//_"github.com/go-sql-driver/mysql"
	//_"mysqlserver"
	"database/sql"
	_"net/smtp"
  //_"uuid"
	"time"
	"os"
	"io"
	"strings"
	"crypto/sha1"
	"io/ioutil"
	"github.com/gorilla/sessions"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	//"github.com/kabukky/httpscerts"
	"strconv"
  "path/filepath"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/ChaitanyaAkula/gittyjobs0"
	"github.com/ChaitanyaAkula/gittyjobs1"
	"github.com/ChaitanyaAkula/gittyjobs2"
	//"github.com/ChaitanyaAkula/gittyjobs3"
	"github.com/ChaitanyaAkula/gittyjobs4"
	"github.com/ChaitanyaAkula/gittyjobs5"
	//"cookies"
	//"imageupload"
	"mime/multipart"
	//"sendotp"
	//"gopkg.in/gomail.v2"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)
func companyregis(w http.ResponseWriter,req *http.Request){
	tpl.ExecuteTemplate(w,"companyregistration.tmpl"," ")
}
func validatecompanyregis(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	Cname:=req.FormValue("cname")
	Cemail:=req.FormValue("email")
	Cpassword:=req.FormValue("pwd")
   chaitu:=[]string{Cemail}

   rows,_:=db.Query("select email from companylogin WHERE email=?",Cemail)
   for rows.Next(){
	   err:=rows.Scan(&Companyemail)
		   
	   if err!=nil{
		   log.Fatalln(err)
	   }
 }	
 if Companyemail==Cemail{
	tpl.ExecuteTemplate(w,"companyregistration.tmpl","email is already existed")
 } else{
	tpl.ExecuteTemplate(w,"companyemailverification.tmpl",chaitu)
	
 }
 

}
func companyemailverification2(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	 defer db.Close()
	
	Cname:=req.FormValue("cname")
	Cemail:=req.FormValue("cemail")
	Cpassword:=req.FormValue("cpassword")
  balance:=-1
  cookies.SetCookie(w,"Cname",Cname)
  cookies.SetCookie(w,"Cemail",Cemail)
  cookies.SetCookie(w,"Cpassword",Cpassword)
	insForm,err1 := db.Prepare("INSERT INTO companylogin(email,companyname,password,balance,location,website,companysize,founded,industrytype,logo,stripeToken) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
		if err1!=nil{
		  log.Fatalln(err1)
	  }
		_,err:= insForm.Exec(Cemail,Cname,Cpassword,balance,"CompanyLocation","Company Website","Select one", "2019"," "," logo.jpg"," ")
		if err!=nil{
		  log.Fatalln(err)
		}
		
		http.Redirect(w,req,"/editcompanyprofile",307)
}
func editcompanyprofile(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	 defer db.Close()
//	Companyname,_:=req.Cookie("Cname")
	Compemail,_:=req.Cookie("Cemail")
	Comppassword,_:=req.Cookie("Cpassword")

	rows,_:=db.Query("select idcompanylogin,email,companyname,location,website,companysize,founded,logo,industrytype from companylogin WHERE email=? and password=?",Compemail.Value,Comppassword.Value)
	for rows.Next(){
		err:=rows.Scan(&Companyid,&Companyemail,&Companyname,&Companylocation,&Companywebsite,&Companysize,&Companyfounded,&Companylogo,&Companytype)
			
		if err!=nil{
			log.Fatalln(err)
		}
  }	
  
  cookies.SetCookie(w,"Companyid",Companyid)
  chaitu:=[]string{Companyname,Companylocation,Companyemail,Companywebsite,Companysize,Companyfounded,Companylogo,Companytype}
	tpl.ExecuteTemplate(w,"companyEditProfile.tmpl",chaitu)
}
func savecompanyprofile(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()  
	Comppassword,_:=req.Cookie("Cpassword")
	fmt.Println(Comppassword.Value)
	cookies.SetCookie(w,"Cemail",req.FormValue("cemail"))
    
    Companylocation=req.FormValue("location")
	Companywebsite=req.FormValue("cwebsite")
	Companysize=req.FormValue("size")
	Companyfounded=req.FormValue("founded")
	
    Companytype=req.FormValue("industry")
	mf, fh, _ := req.FormFile("profile")
		

	Comid,_:=req.Cookie("Companyid")
	//mf, fh, _ := req.FormFile("logo")
			
   //fname:=imageupload.SetImage(w,mf,fh)
   if mf!=nil{
	Companylogo=imageupload.SetImage(w,mf,fh)
	insForm,err1 := db.Prepare("UPDATE companylogin SET logo=? WHERE idcompanylogin=? ")
	if err1!=nil{
		 log.Fatalln(err1)
		}
		_,err:= insForm.Exec(Companylogo,Comid.Value)
	 if err!=nil{
						log.Fatalln(err)
			}
   }
   if Companylocation!=""{
	insForm,err1 := db.Prepare("UPDATE companylogin SET location= ? WHERE idcompanylogin=? ")
	if err1!=nil{
		 log.Fatalln(err1)
		}
		_,err:= insForm.Exec(Companylocation,Comid.Value)
	 if err!=nil{
						log.Fatalln(err)
			}
   }
   if Companysize!=""{
	insForm,err1 := db.Prepare("UPDATE companylogin SET companysize= ? WHERE idcompanylogin=? ")
	if err1!=nil{
		 log.Fatalln(err1)
		}
		_,err:= insForm.Exec(Companysize,Comid.Value)
	 if err!=nil{
						log.Fatalln(err)
			}
   }
    if Companyfounded!=""{
	insForm,err1 := db.Prepare("UPDATE companylogin SET founded= ? WHERE idcompanylogin=? ")
	if err1!=nil{
		 log.Fatalln(err1)
		}
		_,err:= insForm.Exec(Companyfounded,Comid.Value)
	 if err!=nil{
						log.Fatalln(err)
			}
   }
    if Companytype!=""{
	insForm,err1 := db.Prepare("UPDATE companylogin SET industrytype= ? WHERE idcompanylogin=? ")
	if err1!=nil{
		 log.Fatalln(err1)
		}
		_,err:= insForm.Exec(Companytype,Comid.Value)
	 if err!=nil{
						log.Fatalln(err)
			}

  } 
 
 if  Companytype!="" && Companysize!=""&& Companyfounded!=""&& Companylocation!=""{
	insForm,err1 := db.Prepare("UPDATE companylogin SET location= ?,website=?,companysize=?,founded=?,industrytype=? WHERE idcompanylogin=? ")
	if err1!=nil{
		 log.Fatalln(err1)
		}
		_,err:= insForm.Exec(Companylocation,Companywebsite,Companysize,Companyfounded,Companytype,Comid.Value)
	 if err!=nil{
						log.Fatalln(err)
			}
}

		http.Redirect(w,req,"/editcompanyprofile",307)
}

func companysignin(w http.ResponseWriter,req *http.Request){
   msg:=" "
	tpl.ExecuteTemplate(w,"companylogin.tmpl",msg)
}
func validatecompanysignin(w http.ResponseWriter,req *http.Request){
	   
		cookies.SetCookie(w,"Cemail",req.FormValue("email"))
		cookies.SetCookie(w,"Cpassword",req.FormValue("pwd"))
		cookies.SetCookie(w,"jobtitle"," ")
		cookies.SetCookie(w,"companylocation"," ")
		cookies.SetCookie(w,"aboutcompany"," ")
		cookies.SetCookie(w,"companysize",".....")
		cookies.SetCookie(w,"jobtype","choose one")
		cookies.SetCookie(w,"hires"," ")
		cookies.SetCookie(w,"jobdescription"," ")
		cookies.SetCookie(w,"companytype"," ")
		cookies.SetCookie(w,"salary"," ")
		cookies.SetCookie(w,"salarytypes"," ")
		cookies.SetCookie(w,"companywebsite"," ")
		cookies.SetCookie(w,"companyemail"," ")
     
     http.Redirect(w,req,"/validatecompanysignin2",307)
}
func validatecompanysignin2(w http.ResponseWriter,req *http.Request){
	
	 db:=dbconnection.Connection()
	 defer db.Close()
	if req.Method == "GET" {
    
		Emailid,_:=req.Cookie("Cemail")
		Password,_:=req.Cookie("Cpassword")
    rows,_:=db.Query("select idcompanylogin,email,companyname,password,balance,location,website,companysize,industrytype,logo from companylogin WHERE email=? and password=?",Emailid.Value,Password.Value)
			for rows.Next(){
				err:=rows.Scan(&Companyid,&Companyemail,&Companyname,&Companypassword,&Companybalance,&Companylocation,&Companywebsite,&Companysize,&Companytype,&Companylogo)
					
				if err!=nil{
					log.Fatalln(err)
				}
		  }	
	
			if Companyemail==Emailid.Value && Companypassword==Password.Value{
			  
				session, _ := store.Get(req, "session")
				session.Values["authenticated"] = true
				session.Save(req,w)
				cookies.SetCookie(w,"Companyid",Companyid)
				cookies.SetCookie(w,"companyname",Companyname)
				cookies.SetCookie(w,"companyid",Companyid)
				cookies.SetCookie(w,"companylocation",Companylocation)
				cookies.SetCookie(w,"companywebsite",Companywebsite)
				cookies.SetCookie(w,"companysize",Companysize)
				cookies.SetCookie(w,"companylogo",Companylogo)
				cookies.SetCookie(w,"companytype",Companytype)
				cookies.SetCookie(w,"companyemail",Companyemail)
				
				http.Redirect(w,req,"/companyprofile",307)		
			}else{
				errormsg:="email or password is incorrect"
				tpl.ExecuteTemplate(w,"companylogin.tmpl",errormsg)
			}
	
		}
}
func companyprofile(w http.ResponseWriter,req *http.Request){
	
	 db:=dbconnection.Connection()
	defer db.Close()
	var details JobDetails3
	
	//var compro companystruct
	comid,_:=req.Cookie("Companyid")

var jobdetails4 []JobDetails3=nil

	fmt.Println(jobdetails)
	
	rows1,_:=db.Query("select Companyuniqueid,jobtitle,location,aboutcompany,jobtype,hires,jobdescription,industry,salary,salarytypes,postedon,status from company WHERE idcompany=?",comid.Value)
	for rows1.Next(){
		err:=rows1.Scan(&Companyid,&Jobtitle,&Companylocation,&Aboutcompany,&Jobtype,&Hires,&Jobdesc,&Industry,&Salary,&Salarytypes,&Companypostedon,&Companystatus)
		if err!=nil{
			log.Fatalln(err)
		}
		details.Companyid=Companyid
		details.Jobtitle=Jobtitle
		details.Companylocation =Companylocation
		details.Aboutcompany=Aboutcompany
		details.Jobtype=Jobtype
		details.Hires=Hires
		details.Jobdescription=Jobdesc
		details.Industry=Industry
		details.Salary=Salary
		details.Salarytypes=Salarytypes	
        details.Postedon=Companypostedon
		details.Status=Companystatus	
	     jobdetails4=append(jobdetails4,details)
		
		}	
	
		
	
  tpl.ExecuteTemplate(w,"companyprofile.tmpl",jobdetails4)



}
func openJobs(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	var details JobDetails3
	
	//var compro companystruct
	comid,_:=req.Cookie("Companyid")

     var jobdetails4 []JobDetails3=nil

	fmt.Println(jobdetails)
	
	rows1,_:=db.Query("select Companyuniqueid,jobtitle,location,aboutcompany,jobtype,hires,jobdescription,industry,salary,salarytypes from company WHERE idcompany=? and status=?",comid.Value,"open")
	for rows1.Next(){
		err:=rows1.Scan(&Companyid,&Jobtitle,&Companylocation,&Aboutcompany,&Jobtype,&Hires,&Jobdesc,&Industry,&Salary,&Salarytypes)
		if err!=nil{
			log.Fatalln(err)
		}
		details.Companyid=Companyid
		details.Jobtitle=Jobtitle
		details.Companylocation =Companylocation
		details.Aboutcompany=Aboutcompany
		details.Jobtype=Jobtype
		details.Hires=Hires
		details.Jobdescription=Jobdesc
		details.Industry=Industry
		details.Salary=Salary
		details.Salarytypes=Salarytypes	

	     jobdetails4=append(jobdetails4,details)
		
		}	
	tpl.ExecuteTemplate(w,"openJobs.tmpl", jobdetails4)
  }
  func closedJobs(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	var details JobDetails3
	
	//var compro companystruct
	comid,_:=req.Cookie("Companyid")

     var jobdetails4 []JobDetails3=nil

	fmt.Println(jobdetails)
	
	rows1,_:=db.Query("select Companyuniqueid,jobtitle,location,aboutcompany,jobtype,hires,jobdescription,industry,salary,salarytypes from company WHERE idcompany=? and status=?",comid.Value,"close")
	for rows1.Next(){
		err:=rows1.Scan(&Companyid,&Jobtitle,&Companylocation,&Aboutcompany,&Jobtype,&Hires,&Jobdesc,&Industry,&Salary,&Salarytypes)
		if err!=nil{
			log.Fatalln(err)
		}
		details.Companyid=Companyid
		details.Jobtitle=Jobtitle
		details.Companylocation =Companylocation
		details.Aboutcompany=Aboutcompany
		details.Jobtype=Jobtype
		details.Hires=Hires
		details.Jobdescription=Jobdesc
		details.Industry=Industry
		details.Salary=Salary
		details.Salarytypes=Salarytypes	

	     jobdetails4=append(jobdetails4,details)
		
		}	  
	tpl.ExecuteTemplate(w,"closedJobs.tmpl", jobdetails4)
	}
	func updatejobstatus(w http.ResponseWriter,req *http.Request){
		db:=dbconnection.Connection()
	     defer db.Close()
	 comid:= req.FormValue("bookId")
	 insForm,err1 := db.Prepare("UPDATE company SET status= ? WHERE companyuniqueid=? ")
      if err1!=nil{
	          log.Fatalln(err1)
	  }
	  _,err:= insForm.Exec("close",comid)
      if err!=nil{
					log.Fatalln(err)
		}

     	http.Redirect(w,req,"/companyprofile",307)
	}
func postajob0(w http.ResponseWriter,req *http.Request){
	cookies.SetCookie(w,"jobtitle"," ")
	cookies.SetCookie(w,"companylocation"," ")
	cookies.SetCookie(w,"aboutcompany"," ")
	cookies.SetCookie(w,"companysize",".....")
	cookies.SetCookie(w,"jobtype","choose one")
	cookies.SetCookie(w,"hires"," ")
	cookies.SetCookie(w,"jobdescription"," ")
	cookies.SetCookie(w,"companytype"," ")
	cookies.SetCookie(w,"salary"," ")
	cookies.SetCookie(w,"salarytypes"," ")
	cookies.SetCookie(w,"companywebsite"," ")
	cookies.SetCookie(w,"companyemail"," ")
	cookies.SetCookie(w,"skill1"," ")
	cookies.SetCookie(w,"skill2"," ")
	cookies.SetCookie(w,"skill3"," ")
	cookies.SetCookie(w,"otherskills"," ")
	cookies.SetCookie(w,"tools"," ")
	cookies.SetCookie(w,"minexp"," ")
	cookies.SetCookie(w,"maxexp"," ")
	fmt.Println("tested")
 http.Redirect(w,req,"/postajob",307)
}
func postajob(w http.ResponseWriter,req *http.Request){

	
	db:=dbconnection.Connection()
	defer db.Close()
	comid,_:=req.Cookie("Companyid")
    rows,_:=db.Query("select companyname,location,companysize,logo,industrytype,founded,website from companylogin WHERE idcompanylogin=?",comid.Value)
			for rows.Next(){
				err:=rows.Scan(&Companyname,&Companylocation,&Companysize,&Companylogo,&Companytype,&Companyfounded,&Companywebsite)
					
				if err!=nil{
					log.Fatalln(err)
				}
		  }				
		    jobtitle,_:=req.Cookie("jobtitle")
			aboutcompany,_:=req.Cookie("aboutcompany")
		     //jobtype,_:=req.Cookie("jobtype")
			
				companypro:=[]string{Companyname,jobtitle.Value,Companylocation,aboutcompany.Value,Companysize,Companylogo,Companytype,Companyfounded,Companywebsite}
			    err:= tpl.ExecuteTemplate(w,"jobposting1.tmpl",companypro)
				if err!=nil{
					log.Fatalln(err)
			   }	
	
	
}
func postedjobs(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	var details JobDetails
	comid,_:=req.Cookie("companyid")
	var jobdetails []JobDetails=nil
	fmt.Println(jobdetails)
	rows,_:=db.Query("select companyname,jobtitle,location,salary,salarytypes from company WHERE idcompany=?",comid.Value)
	for rows.Next(){
		err:=rows.Scan(&Companyname,&Jobtitle,&Companylocation,&Salary,&Salarytypes)
		if err!=nil{
			log.Fatalln(err)
		}
		details.Companyname=Companyname
		details.Jobtitle=Jobtitle
		details.Location =Companylocation
		details.Salary=Salary
		details.Salarytypes=Salarytypes	
		
		jobdetails=append(jobdetails,details)
		
		fmt.Println(jobdetails)
		
		}	
	//jobdetails:=[]string{Companyname,Jobtitle,Companylocation,Salary,Salarytypes}	
	
	tpl.ExecuteTemplate(w,"postedjobs.tmpl",jobdetails)
	 
	
}
func aboutcandidate(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	companyid:=req.FormValue("companyid")
	var candi candidatestruct
//	var candi1 candidatestruct
	var candidatedetails []candidatestruct=nil
//	var candidatedetails1 []candidatestruct=nil
/*rows1,_:=db.Query("select jobtitle from company WHERE companyuniqueid=?",companyid)
for rows1.Next(){
	err:=rows1.Scan(&Jobtitle)
	if err!=nil{
		log.Fatalln(err)
	}
	
	candi1.Jobtitle=Jobtitle

}		

candidatedetails=append(candidatedetails,candi1)*/

    rows,_:=db.Query("select iddeveloper from candidate WHERE companyuniqueid=?",companyid)
			for rows.Next(){
				err:=rows.Scan(&developerid)
				if err!=nil{
					log.Fatalln(err)
				}
				rows1,_:=db.Query("select iddeveloper,firstname,lastname from developer WHERE iddeveloper=?",developerid)
				for rows1.Next(){
					err:=rows1.Scan(&developerid,&developerFirstname,&developerLastname)
					if err!=nil{
						log.Fatalln(err)
					}
					
			  }			
				 candi.Devfirstname=developerFirstname
				 candi.Devlastname=developerLastname
				 candi.Devid=developerid
				 candidatedetails=append(candidatedetails,candi)
			}
  
			
				
fmt.Println( candidatedetails)
  tpl.ExecuteTemplate(w,"candidates.tmpl",candidatedetails)
}
func candidateprofile(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	devid:=req.FormValue("devid")
	rows1,_:=db.Query("select firstname,lastname,email,mobile from developer WHERE iddeveloper=?",devid)
       for rows1.Next(){
	          err:=rows1.Scan(&developerFirstname,&developerLastname,&developerEmail,&developerMobile)
	         if err!=nil{
		            log.Fatalln(err)
	         }
			}
			rows2,_:=db.Query("select imagepath,location,resumepath from profile WHERE iddeveloper=?",devid)
			for rows2.Next(){
				   err:=rows2.Scan(&developerImage,&developerLocation,&developerResumepath)
				  if err!=nil{
						 log.Fatalln(err)
				  }
				 }
		    rows3,_:=db.Query("select designation,year,month from devqualification WHERE iddeveloper=?",devid)
			 for rows3.Next(){
					err:=rows3.Scan(&devqualDesignation,&devqualYear,&devqualMonth)
					if err!=nil{
						  log.Fatalln(err)
					 }
			 }
			 rows4,_:=db.Query("select skill1,skill2,skill3 from devskills WHERE iddeveloper=?",devid)
			 for rows4.Next(){
					err:=rows4.Scan(&developerSkill1,&developerSkill2,&developerSkill3)
					if err!=nil{
						  log.Fatalln(err)
					 }
			 }
	  candi:=[]string{developerFirstname,developerLastname,developerEmail,developerMobile,developerImage,developerLocation,developerResumepath,devqualDesignation,devqualYear,devqualMonth,developerSkill1,developerSkill2,developerSkill3}
	  
		
	tpl.ExecuteTemplate(w,"candidateProfile.tmpl",candi)
  }
func aboutpostingjob(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close() 
	jobid:=req.FormValue("jobid")
	rows1,_:=db.Query("select idcompany,jobtitle,companyname,location,website,size,industry,companyemail,aboutcompany,jobtype,hires,jobdescription,salary,salarytypes,firstskill,secondskill,thirdskill,otherskill,toolsandtechnologies,minimumexp,maximumexp from company WHERE companyuniqueid=?",jobid)
	for rows1.Next(){
		err:=rows1.Scan(&Companyid,&Jobtitle,&Companyname,&Companylocation,&Companywebsite,&Companysize,&Companytype,&Companyemail,&Aboutcompany,&Jobtype,&Hires,&Jobdesc,&Salary,&Salarytypes,&Companyskill1,&Companyskill2,&Companyskill3,&Companyotherskills,&Companytools,&Companyminexp,&Companymaxexp)
		if err!=nil{
			log.Fatalln(err)
		}
	}
	rows2,_:=db.Query("select logo from companylogin WHERE idcompanylogin=?",Companyid)
	for rows2.Next(){
		err:=rows2.Scan(&Companylogo)
		if err!=nil{
			log.Fatalln(err)
		}
	}
	com:=[]string{Jobtitle,Companyname,Companylocation,Companywebsite,Companysize,Companytype,Companyemail,Companylogo,Aboutcompany,Jobtype,Hires,Jobdesc,Salary,Salarytypes,Companyskill1,Companyskill2,Companyskill3,Companyotherskills,Companytools,Companyminexp,Companymaxexp}

	tpl.ExecuteTemplate(w,"aboutpostingjob.tmpl",com)
  }
func companyprofile2(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	comid,_:=req.Cookie("Companyid")
    rows,_:=db.Query("select companyname,website,email from companylogin WHERE idcompanylogin=?",comid.Value)
			for rows.Next(){
				err:=rows.Scan(&Companyname,&Companywebsite,&Companyemail)
					
				if err!=nil{
					log.Fatalln(err)
				}
		  }				
	cookies.SetCookie(w,"companyname",Companyname)
	cookies.SetCookie(w,"companywebsite",Companywebsite)
	cookies.SetCookie(w,"companyemail",Companyemail)
	cookies.SetCookie(w,"jobtitle",req.FormValue("jobtitle"))
	cookies.SetCookie(w,"companylocation",req.FormValue("location"))
	cookies.SetCookie(w,"aboutcompany",req.FormValue("aboutcompany"))
	cookies.SetCookie(w,"companysize", req.FormValue("size"))
	
	

 http.Redirect(w,req,"/companyprofile2a",307)
}
func companyprofile2a(w http.ResponseWriter,req *http.Request){
	comname,_:=req.Cookie("companyname")
	jobtitle,_:=req.Cookie("jobtitle")
	hires,_:=req.Cookie("hires")
  jobdescription,_:=req.Cookie("jobdescription")
  jobtype,_:=req.Cookie("jobtype")
	salary,_:=req.Cookie("salary")
	salarytypes,_:=req.Cookie("salarytypes")
	website,_:=req.Cookie("companywebsite")
	companyemail,_:=req.Cookie("companyemail")
	companypro:=[]string{comname.Value,jobtitle.Value,hires.Value,jobdescription.Value,jobtype.Value,salary.Value,salarytypes.Value,website.Value,companyemail.Value}
	err:= tpl.ExecuteTemplate(w,"jobposting2.tmpl",companypro)
	if err!=nil{
		log.Fatalln(err)
	 }	
}
func companyaddskills(w http.ResponseWriter,req *http.Request){
	cookies.SetCookie(w,"hires",req.FormValue("hires"))
	jobdesc:=req.FormValue("jobdescription")
	fmt.Println(jobdesc)
	cookies.SetCookie(w,"jobdescription",jobdesc)
  cookies.SetCookie(w,"jobtype",req.FormValue("jobtype"))
	cookies.SetCookie(w,"salary",req.FormValue("salary"))
	cookies.SetCookie(w,"salarytypes",req.FormValue("salarytypes"))
	cookies.SetCookie(w,"companywebsite",req.FormValue("website"))
	cookies.SetCookie(w,"companyemail",req.FormValue("companyemail"))
	jobtitle,_:=req.Cookie("jobtitle")
	skill1,_:=req.Cookie("skill1")
	skill2,_:=req.Cookie("skill2")
	skill3,_:=req.Cookie("skill3")
	otherskills,_:=req.Cookie("otherskills")
	tools,_:=req.Cookie("tools")
	minexp,_:=req.Cookie("minexp")
	maxexp,_:=req.Cookie("maxexp")
	com:=[]string{jobtitle.Value,skill1.Value,skill2.Value,skill3.Value,otherskills.Value,tools.Value,minexp.Value,maxexp.Value}
	tpl.ExecuteTemplate(w,"companyaddSkills.tmpl",com)

}
func companyprofile3(w http.ResponseWriter,req *http.Request){
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
		 cookies.SetCookie(w,"skill1",req.FormValue("skill1"))
		 cookies.SetCookie(w,"skill2",req.FormValue("skill2"))
	   cookies.SetCookie(w,"skill3",req.FormValue("skill3"))
		 cookies.SetCookie(w,"otherskills",req.FormValue("otherskills"))
		 cookies.SetCookie(w,"tools",req.FormValue("tools"))
		 cookies.SetCookie(w,"minexp",req.FormValue("minexp"))
		 cookies.SetCookie(w,"maxexp",req.FormValue("maxexp"))
  err:= tpl.ExecuteTemplate(w,"jobposting3.tmpl",Companybalance)
	if err!=nil{
		log.Fatalln(err)
	 }	
 
}