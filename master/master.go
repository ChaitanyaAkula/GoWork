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
var (
	googleOauthConfig *oauth2.Config
	oauthStateString = "pseudo-random"
)
var Applytest bool
var userprofile []string
var imagepath,language,iddev string 
var developerid,developerFirstname,developerLastname,developerEmail,developerMobile,developerLocation,developerResumepath,developerPassword,developerImage string
var developerDegree,developerSpecialization,developerInstitutename,developerType,developerPassingyear string
var developerSkill1,developerSkill2,developerSkill3 string
var devqualCompanyname,devqualDesignation,devqualLocation,devqualStartdate,devqualEnddate,devqualDescription,devqualYear,devqualMonth,devqualexperience string
var Companyid,Companyname,Companyemail,Companypassword,Companylocation,Companywebsite,Companysize,Companylogo,Companytype,Companyfounded string
var Jobtitle,Aboutcompany,Jobtype,Hires,Industry,Jobdesc,Salary,Salarytypes,Stripetoken string
var Companyskill1,Companyskill2,Companyskill3,Companyotherskills,Companytools,Companyminexp,Companymaxexp,Companypostedon,Companystatus string
var Count_fulltime,Count_parttime,Count_internship string
var mf multipart.File
var fh *multipart.FileHeader
var Companybalance int
var jobdetails =make([]JobDetails,0)
var jobDetails3 =make([]JobDetails3,0)
var candidatedetails =make([]candidatestruct,0)
var devfname,devlname,devemail,devmobile string
var db *sql.DB
var tpl *template.Template
var store = sessions.NewCookieStore([]byte("super-secret-key"))
type adminlogin struct{
	email string
	password string
}
type devprofilestruct struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}
type developer struct{
	fname string
	lname string
	email string
	password string
	mobile string
}
type JsonType struct{
	Array string
}

type companylogin struct{
	Companyname string
	Email string
}
type person struct{
	//Id string
	Location string
	JobTitle string
	Email string
	JobDescription string
	Company string
	Website string
	Apply string
}
type otp struct{
	Mobile string
	Msgg int64
}
type languagestruct struct{
	lang []string 
}
type company struct{
	//Id string
	Location string
	JobTitle string
	Email string
	JobDescription string
	Company string
	Stripe string
	
}
type JobDetails struct{
	Companyname string
	Jobtitle string
	Location string
	Salary string
	Salarytypes string
//	Details []string 
}
type JobDetails3 struct{

	Companyid string
	Companyname string
	Companyemail string
	Companylocation string
	Companywebsite string
	Companylogo string
	Jobtitle string
	Aboutcompany string
	Jobtype string
	Hires string
	Jobdescription string
	Industry string
	Salary string
	Salarytypes string
	Postedon string
	Status string
//	Details []string 
}
type companystruct struct{
	details0 []JobDetails3
	details1 []JobDetails3
}
type JobDetails1 struct{

	Companyid string
	Companyname string
	Jobtitle string
	Location string
	Salary string
	Salarytypes string
	Jobtype string 
	Hires string
	Jobdesc string
	Industry string
	Companywebsite string
	Companyemail string
	Companyskill1 string
	Companyskill2 string
	Companyskill3 string
	Companyotherskills string
	Companytools string
	Companyminexp string
	Companymaxexp string
	Companypostedon string
	Count_fulltime string
	Count_parttime string
	Count_internship string
}
type profile struct{
	id string `json:"id"`
	email string `json:"email"`
	verified_email bool `json:"verified_email"`
	picture string `json:"picture"`

}
type profiles []profile

type userstruct struct {
	 id string	
	 Firstname string
	  Lastname    string
		Email    string
		Mobile string
	//	Path  string
}
type candidatestruct struct{
	Devfirstname string
	Devlastname string
	Devid string
}
var m,_ = time.ParseDuration("10m")
var expiration = time.Now().Add(m)

func init(){


	tpl = template.Must(template.New("").ParseGlob("templates/*.tmpl"))
	
}
func main(){

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	myRouter.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	myRouter.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	myRouter.HandleFunc("/",Home)
	myRouter.HandleFunc("/searchjobs",searchjobs)
	myRouter.HandleFunc("/jobresults",jobresults)
	myRouter.HandleFunc("/login",login).Methods("GET")
	myRouter.HandleFunc("/signin",devsignin).Methods("GET")
	myRouter.HandleFunc("/companyregis",companyregis).Methods("GET")
	myRouter.HandleFunc("/validatecompanyregis",validatecompanyregis).Methods("GET")
	myRouter.HandleFunc("/companysignin",companysignin).Methods("GET")
	myRouter.HandleFunc("/validatecompanysignin",validatecompanysignin).Methods("GET")
	myRouter.HandleFunc("/validatecompanysignin2",validatecompanysignin2).Methods("GET")
	myRouter.HandleFunc("/companyprofile",companyprofile)
	myRouter.HandleFunc("/postajob",postajob)
	myRouter.HandleFunc("/postajob0",postajob0)
	myRouter.HandleFunc("/postedjobs",postedjobs)
	myRouter.HandleFunc("/openJobs",openJobs)
	myRouter.HandleFunc("/updatejobstatus",updatejobstatus)
    myRouter.HandleFunc("/closedJobs",closedJobs)
	myRouter.HandleFunc("/aboutcandidate",aboutcandidate)
	myRouter.HandleFunc("/candidateprofile",candidateprofile)
	myRouter.HandleFunc("/aboutpostingjob",aboutpostingjob)
	myRouter.HandleFunc("/companyprofile2",companyprofile2)
	myRouter.HandleFunc("/companyprofile2a",companyprofile2a)
	myRouter.HandleFunc("/companyaddskills",companyaddskills)
	myRouter.HandleFunc("/companyprofile3",companyprofile3)
	myRouter.HandleFunc("/companyemailverification2",companyemailverification2).Methods("GET")
	myRouter.HandleFunc("/editcompanyprofile",editcompanyprofile)
	myRouter.HandleFunc("/savecompanyprofile",savecompanyprofile)
	myRouter.HandleFunc("/savecompanylogo",savecompanyprofile)
	myRouter.HandleFunc("/devsignin2",devsignin2).Methods("GET")
	myRouter.HandleFunc("/devlogout",devlogout).Methods("GET")
	myRouter.HandleFunc("/validatesignin",validatedevsignin).Methods("GET")
   myRouter.HandleFunc("/verify",verify)
   myRouter.HandleFunc("/verify1",verify1)
	myRouter.HandleFunc("/devprofile",devprofile)
	myRouter.HandleFunc("/editprofile",editprofile)
	myRouter.HandleFunc("/saveprofilepic",saveprofilepic).Methods("POST")
	myRouter.HandleFunc("/savelanguage",savelanguage).Methods("POST")
	myRouter.HandleFunc("/uploadpic",uploadpic)
	myRouter.HandleFunc("/signup",devsignup)
	myRouter.HandleFunc("/verifyaccount",verifyaccount)
	myRouter.HandleFunc("/otpvalidation",otpvalidation)
	myRouter.HandleFunc("/stripePayment1",stripePayment1).Methods("POST")
	myRouter.HandleFunc("/stripePayment2",stripePayment2).Methods("POST")
	myRouter.HandleFunc("/stripePayment3",stripePayment3).Methods("POST")
	myRouter.HandleFunc("/post",post)
	myRouter.HandleFunc("/postconfirmation",postconfirmation)
	myRouter.HandleFunc("/adminlogingittyapps",admin).Methods("GET")
	myRouter.HandleFunc("/validateadmin",validateadmin).Methods("GET")
	myRouter.HandleFunc("/validatelogin",validatelogin).Methods("GET")
	myRouter.HandleFunc("/sendmail",sendmail)
	myRouter.HandleFunc("/googlesignin",googlesignin)
	myRouter.HandleFunc("/profile",profilepage)
	myRouter.HandleFunc("/newpassword",newpassword)
	myRouter.HandleFunc("/confirmpassword",confirmpassword)
	myRouter.HandleFunc("/editprofile2",editprofile2)
	myRouter.Handle("/favicon.ico",http.NotFoundHandler())
	myRouter.HandleFunc("/freejobalert",freejobalert)
	myRouter.HandleFunc("/freejobalert/subscribed",freejobalertsubscribed)
	myRouter.HandleFunc("/about_GittyJobs",about_GittyJobs)
	myRouter.HandleFunc("/aboutGitty",aboutGitty)
	myRouter.HandleFunc("/contactus",contact_GittyJobs)
	myRouter.HandleFunc("/contactus.",contact_GittyJobs1)
	myRouter.HandleFunc("/carrer",carrer)
	myRouter.HandleFunc("/ourservices",ourservices)
	myRouter.HandleFunc("/postfreetrial",postfreetrial)
	////////////////////developer add profile///////////
	myRouter.HandleFunc("/devaddprofile/personaldetails",personaldetails)
	myRouter.HandleFunc("/devpersonaldetails00",devpersonaldetails00)
	myRouter.HandleFunc("/devaddprofile/educationaldetails",educationaldetails)
	myRouter.HandleFunc("/deveducationaldetails00",deveducationaldetails00)
	myRouter.HandleFunc("/devaddprofile/qualificationdetails",qualificationdetails)
	myRouter.HandleFunc("/qualificationdetails00",qualificationdetails00)
	///////////registering filters////////////
	myRouter.HandleFunc("/searchjobs0",searchjobs0).Methods("GET")
	myRouter.HandleFunc("/filter_fulltime",filter_fulltime).Methods("GET")
	myRouter.HandleFunc("/filter_parttime",filter_parttime).Methods("GET")
	myRouter.HandleFunc("/filter_internship",filter_internship).Methods("GET")




	//////////end////////////////////

	///////////////Apply Job/////////
myRouter.HandleFunc("/applyjob/signin",applyjobsignin)
myRouter.HandleFunc("/applyjob/applied",applyjobapplied)

	//////////////end///////////////

	http.ListenAndServe(":8080",myRouter)
	
}
func Home(w http.ResponseWriter, req *http.Request ){
	tpl.ExecuteTemplate(w, "index.tmpl", nil)
}
func devsignin(w http.ResponseWriter, req *http.Request ){
	Applytest=true
	tpl.ExecuteTemplate(w, "signin.tmpl", " ")
}
func devsignin2(w http.ResponseWriter, req *http.Request ){
	m,_ := time.ParseDuration("1m")
			expire := time.Now().Add(m)  
			
	http.SetCookie(w,&http.Cookie{
		Name:"email",
		Value: req.FormValue("email"),
		Expires:expire,
		HttpOnly: true,
	})
	http.SetCookie(w,&http.Cookie{
		Name:"password",
		Value: req.FormValue("password"),
		Expires:expire,
		HttpOnly: true,
	})
	
	http.Redirect(w,req,"/validatesignin",307)
	

}
func validatedevsignin (w http.ResponseWriter, req *http.Request ){
	db:=dbconnection.Connection()
	 defer db.Close()
    

		
		Emailid,_:=req.Cookie("email")
		Password,_:=req.Cookie("password")
		
		rows,_:=db.Query("select iddeveloper,email,psw from developer WHERE email=? and psw=?",Emailid.Value,Password.Value)
			for rows.Next(){
				err:=rows.Scan(&developerid,&developerEmail,&developerPassword)
				if err!=nil{
					log.Fatalln(err)
				}
				}	
				http.SetCookie(w,&http.Cookie{
					Name:"devid",
					Value: developerid,
					
				})
			
		
			if developerEmail==Emailid.Value && developerPassword==Password.Value{
			  
				session, _ := store.Get(req, "session")
				session.Values["authenticated"] = true
				session.Save(req,w)
				if Applytest== false{
					http.Redirect(w,req,"/applyjob/applied",307)
				}else{
					http.Redirect(w,req,"/devprofile",307)
				}
				
			   
			}else{
				tpl.ExecuteTemplate(w, "signin.tmpl", "email or password is invalid ")
			}
			
	
		

	}
	
func devprofile(w http.ResponseWriter,req *http.Request){
  
	db:=dbconnection.Connection()
	defer db.Close()
	c3,_:=req.Cookie("devid")
	rows,_:=db.Query("select firstname,lastname,email,mobile from developer WHERE iddeveloper=?",c3.Value)
	for rows.Next(){
		err:=rows.Scan(&developerFirstname,&developerLastname,&developerEmail,&developerMobile)
			
		if err!=nil{
			log.Fatalln(err)
		}
	}
	dev:=[]string{developerFirstname,developerLastname,developerEmail,developerMobile}

		tpl.ExecuteTemplate(w,"devprofile.tmpl",dev)
	
		
} 
func editprofile(w http.ResponseWriter,req *http.Request){
	
	db:=dbconnection.Connection()
	defer db.Close()

tpl.ExecuteTemplate(w,"developereditprofile.tmpl",nil)
}
///////////////////////developer add profile///////////////
func personaldetails(w http.ResponseWriter, req *http.Request){
		
	db:=dbconnection.Connection()
	defer db.Close()
	c3,_:=req.Cookie("devid")
	rows,_:=db.Query("select firstname,lastname,email,mobile from developer WHERE iddeveloper=?",c3.Value)
	for rows.Next(){
		err:=rows.Scan(&developerFirstname,&developerLastname,&developerEmail,&developerMobile)
			
		if err!=nil{
			log.Fatalln(err)
		}
	}
	rows,_=db.Query("select location,resumepath from profile WHERE iddeveloper=?",c3.Value)
	for rows.Next(){
		err:=rows.Scan(&developerLocation,&developerResumepath)
			
		if err!=nil{
			log.Fatalln(err)
		}
	}
	
   dev:=[]string{developerFirstname,developerLastname,developerEmail,developerMobile,developerLocation,developerResumepath}
	tpl.ExecuteTemplate(w,"devPersonalDetail.tmpl",dev)
}
func devpersonaldetails00(w http.ResponseWriter,req *http.Request){
    db:=dbconnection.Connection()
	defer db.Close()
   location:=req.FormValue("dlocation")
   mf, fh, _ := req.FormFile("resume")
   c3,_:=req.Cookie("devid")
   if mf!=nil{
	resume:=resumeupload.SetResume(w,mf,fh)
	insForm,err1 := db.Prepare(" UPDATE profile SET resumepath=? WHERE iddeveloper =? ")
	if err1!=nil{
	  log.Fatalln(err1)
  }
	_,err:= insForm.Exec(resume,c3.Value)
	if err!=nil{
	  log.Fatalln(err)
  }
   }
	
   insForm,err1 := db.Prepare(" UPDATE profile SET location=? WHERE iddeveloper =? ")
   if err1!=nil{
	 log.Fatalln(err1)
 }
   _,err:= insForm.Exec(location,c3.Value)
   if err!=nil{
	 log.Fatalln(err)
 }
 http.Redirect(w,req,"/devaddprofile/educationaldetails",307)
}
func educationaldetails(w http.ResponseWriter, req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	c3,_:=req.Cookie("devid")
	rows,_:=db.Query("select degree,specialization,institutename,type,passingyear from deveducationaldetails WHERE iddeveloper=?",c3.Value)
	for rows.Next(){
		err:=rows.Scan(&developerDegree,&developerSpecialization,&developerInstitutename,&developerType,&developerPassingyear)
			
		if err!=nil{
			log.Fatalln(err)
		}
	}
	rows1,_:=db.Query("select skill1,skill2,skill3 from devskills WHERE iddeveloper=?",c3.Value)
	for rows1.Next(){
		err:=rows1.Scan(&developerSkill1,&developerSkill2,&developerSkill3)
			
		if err!=nil{
			log.Fatalln(err)
		}
	}
	fmt.Println(developerSkill3)
	dev:=[]string{developerDegree,developerSpecialization,developerInstitutename,developerType,developerPassingyear,developerSkill1,developerSkill2,developerSkill3}
	tpl.ExecuteTemplate(w,"devEducationalDetail.tmpl",dev)
}
func deveducationaldetails00(w http.ResponseWriter, req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
   degree:=req.FormValue("quali")
   specialization:=req.FormValue("speci")
   insname:=req.FormValue("ins")
   studytype:=req.FormValue("type")
   year:=req.FormValue("year")
   skill1:=req.FormValue("skill1")
   skill2:=req.FormValue("skill2")
   skill3:=req.FormValue("skill3")
   c3,_:=req.Cookie("devid")
   insForm,err1 := db.Prepare(" UPDATE deveducationaldetails SET degree=?,specialization=?,institutename=?,type=?,passingyear=? WHERE iddeveloper =? ")
   if err1!=nil{
	 log.Fatalln(err1)
 }
   _,err:= insForm.Exec(degree,specialization,insname,studytype,year,c3.Value)
   if err!=nil{
	 log.Fatalln(err)
 }
 insForm,err1 = db.Prepare(" UPDATE devskills SET skill1=?,skill2=?,skill3=? WHERE iddeveloper =? ")
 if err1!=nil{
   log.Fatalln(err1)
}
 _,err= insForm.Exec(skill1,skill2,skill3,c3.Value)
 if err!=nil{
   log.Fatalln(err)
}
 http.Redirect(w,req,"/devaddprofile/qualificationdetails",307)
}
func qualificationdetails(w http.ResponseWriter, req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	c3,_:=req.Cookie("devid")
	rows,_:=db.Query("select companyname,designation,location,startdate,enddate,description,year,month,experience from devqualification WHERE iddeveloper=?",c3.Value)
	for rows.Next(){
		err:=rows.Scan(&devqualCompanyname,&devqualDesignation,&devqualLocation,&devqualStartdate,&devqualEnddate,&devqualDescription,&devqualYear,&devqualMonth,&devqualexperience)
			
		if err!=nil{
			log.Fatalln(err)
		}
	}
     dev:=[]string{devqualCompanyname,devqualDesignation,devqualLocation,devqualStartdate,devqualEnddate,devqualDescription,devqualYear,devqualMonth,devqualexperience}
	tpl.ExecuteTemplate(w,"devQualificationDetail.tmpl",dev)
}
func qualificationdetails00(w http.ResponseWriter, req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	c3,_:=req.Cookie("devid")
	devqualexperience:=req.FormValue("radi")
	devqualCompanyname:=req.FormValue("cname")
	devqualDesignation:=req.FormValue("jobtitle")
	devqualStartdate:=req.FormValue("start")
	devqualEnddate:=req.FormValue("end")
	devqualYear:=req.FormValue("year")
	devqualMonth:=req.FormValue("month")
	
	mf, fh, _ := req.FormFile("profile")
	fmt.Println(devqualexperience)
   if mf!=nil{
	devprofile:=devprofileupload.SetProfile(w,mf,fh)
	insForm,err1 := db.Prepare(" UPDATE profile SET imagepath=? WHERE iddeveloper =? ")
	if err1!=nil{
	  log.Fatalln(err1)
  }
	_,err:= insForm.Exec(devprofile,c3.Value)
	if err!=nil{
	  log.Fatalln(err)
  }
   }
   if devqualexperience== "Fresher"{
	   fmt.Println("fresher testing")
   insForm,err1 := db.Prepare(" UPDATE devqualification SET designation=?,experience=? WHERE iddeveloper =? ")
   if err1!=nil{
	 log.Fatalln(err1)
 }
   _,err:= insForm.Exec(devqualDesignation,devqualexperience,c3.Value)
   if err!=nil{
	 log.Fatalln(err)
 }
}
if devqualexperience=="Professional"{
	fmt.Println("pro testing")
	insForm,err1 := db.Prepare(" UPDATE devqualification SET experience=?, designation=?,companyname=?,startdate=?,enddate=?,year=?,month=? WHERE iddeveloper =? ")
	if err1!=nil{
	  log.Fatalln(err1)
  }
	_,err:= insForm.Exec(devqualexperience,devqualDesignation,devqualCompanyname,devqualStartdate,devqualEnddate,devqualYear,devqualMonth ,c3.Value)
	if err!=nil{
	  log.Fatalln(err)
  }
}
http.Redirect(w,req,"/devprofile",307)
}
////////////////////////end///////////////////////////////
func devlogout(w http.ResponseWriter,req *http.Request){
	 m,_ := time.ParseDuration("0m")
      expire := time.Now().Add(m)    
	session, _ := store.Get (req, "session")
	session.Values["authenticated"] = false
	session.Save(req,w)
	http.SetCookie(w,&http.Cookie{
		Name:"email",
		Value: "",
		Expires:expire,
		Path:"/",
		HttpOnly: true,
	})
	http.SetCookie(w,&http.Cookie{
		Name:"password",
		Value: "",
		Expires:expire,
		Path:"/",
		HttpOnly: true,
	})
	
	tpl.ExecuteTemplate(w,"signin.tmpl",nil)
} 

func login(w http.ResponseWriter, req *http.Request ){
	tpl.ExecuteTemplate(w, "login.tmpl", nil)

}
func devsignup(w http.ResponseWriter, req *http.Request ){
	
	err1 :=tpl.ExecuteTemplate(w,"signup.tmpl"," ")
	if err1!=nil{
		log.Fatalln(err1)
	}
	    
}
func verifyaccount(w http.ResponseWriter, req *http.Request ){
	db:=dbconnection.Connection()
	defer db.Close()
	if req.Method == "POST" { 
		Emailid:=req.FormValue("email")
		fname:=req.FormValue("firstname")
		lname:=req.FormValue("lastname")
		mob:=req.FormValue("mobile")
		password:=req.FormValue("psw")
		rows,_:=db.Query("select email from developer WHERE email=?",Emailid)
		for rows.Next(){
			err:=rows.Scan(&developerEmail)
				
			if err!=nil{
				log.Fatalln(err)
			}
	  }	
	  if developerEmail==Emailid{
		tpl.ExecuteTemplate(w,"signup.tmpl","email is already existed ")
	  }else {
                  http.SetCookie(w,&http.Cookie{
			Name:"Emailid",
			Value: Emailid,
		//	Expires:expiration,
		})
		http.SetCookie(w,&http.Cookie{
			Name:"fname",
			Value: fname,
			Expires:expiration,
		})
		http.SetCookie(w,&http.Cookie{
			Name:"lname",
			Value: lname,
		 	Expires:expiration,
		})
		http.SetCookie(w,&http.Cookie{
			Name:"mob",
			Value:mob,
			Expires:expiration,
		})
		http.SetCookie(w,&http.Cookie{
			Name:"password",
			Value: password,
		//	Expires:expiration,
		})
	rand.Seed(time.Now().UnixNano())
	n:=rand.Int63n(100000)
	Msg:=n
	str := strconv.FormatInt(Msg, 10)
	http.SetCookie(w,&http.Cookie{
		Name:"otp",
		Value: str,
		Expires:expiration,
	})
	 

	tpl.ExecuteTemplate(w, "verifyaccount.tmpl", Emailid)
	
   fmt.Println(Msg)
	}
	  }
		
	
}
func otpvalidation(w http.ResponseWriter,req *http.Request){
	otp1:=req.FormValue("otp")
	c,_:=req.Cookie("otp")
	o:=c.Value
	fmt.Println(otp1,c.Value)
	otp, err := strconv.ParseInt(o, 10, 64)
	msg, err := strconv.ParseInt(otp1, 10, 64)
	fmt.Println(otp,msg)
	if err==nil{
		fmt.Println(err)
	}
    if otp==msg{
			
	   http.Redirect(w,req,"/verify",307)
		}
	
		
} 
func verify(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	 defer db.Close()
	
	

	
					
    	c1,_:=req.Cookie("fname")
			 c2,_:=req.Cookie("lname")
			 c3,_:=req.Cookie("Emailid")
			 c4,_:=req.Cookie("mob")
			c5,_:=req.Cookie("password")
	     
			firstname:=c1.Value
			lastname:=c2.Value
			email:=c3.Value
			mobile:=c4.Value
			password:=c5.Value
			fmt.Println(firstname,lastname,email,mobile,password)


			 insForm,err1 := db.Prepare("INSERT INTO developer(firstname,lastname,email,mobile,psw) VALUES(?,?,?,?,?)")
		  if err1!=nil{
			log.Fatalln(err1)
		}
		  _,err:= insForm.Exec(firstname,lastname,email,mobile,password)
		  if err!=nil{
			log.Fatalln(err)
		}
	  
		http.Redirect(w,req,"/verify1",307)
	
}
func verify1(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	 defer db.Close()
	 c3,_:=req.Cookie("Emailid")
	c5,_:=req.Cookie("password")
    rows,_:=db.Query("select iddeveloper from developer WHERE email=? and psw=?",c3.Value,c5.Value)
		 for rows.Next(){
			 err:=rows.Scan(&developerid)
				 
			 if err!=nil{
				 log.Fatalln(err)
			 }
		 }
		 http.SetCookie(w,&http.Cookie{
			Name:"devid",
			Value: developerid,
		
		})
		insForm,err1 := db.Prepare("INSERT INTO profile(iddeveloper,imagepath,location,resumepath) VALUES(?,?,?,?)")
		if err1!=nil{
		  log.Fatalln(err1)
	  }
		_,err:= insForm.Exec(developerid,"image","location","resume")
		if err!=nil{
		  log.Fatalln(err)
	  }
	  insForm,err1 = db.Prepare("INSERT INTO deveducationaldetails(iddeveloper,degree,specialization,institutename,type,passingyear,performancescale,percentage) VALUES(?,?,?,?,?,?,?,?)")
	  if err1!=nil{
		log.Fatalln(err1)
	}
	  _,err= insForm.Exec(developerid,"degree","specialization","institutename","type","passingyear","performancescale","percentage")
	  if err!=nil{
		log.Fatalln(err)
	}
	insForm,err1 = db.Prepare("INSERT INTO devskills(iddeveloper,skill1,skill2,skill3,skill4,skill5,skill6,skill7,skill8,skill9,skill10) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	if err1!=nil{
	  log.Fatalln(err1)
  }
	_,err= insForm.Exec(developerid,"skill1","skill2","skill3","skill4","skill5","skill6","skill7","skill8","skill9","skill10")
	if err!=nil{
	  log.Fatalln(err)
  }
  insForm,err1 = db.Prepare("INSERT INTO devskills(iddeveloper,skill1,skill2,skill3,skill4,skill5,skill6,skill7,skill8,skill9,skill10) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	if err1!=nil{
	  log.Fatalln(err1)
  }
	_,err= insForm.Exec(developerid,"skill1","skill2","skill3","skill4","skill5","skill6","skill7","skill8","skill9","skill10")
	if err!=nil{
	  log.Fatalln(err)
  }
  insForm,err1 = db.Prepare("INSERT INTO devqualification(iddeveloper,companyname,designation,location,startdate,enddate,description,year,month,experience) VALUES(?,?,?,?,?,?,?,?,?,?)")
	if err1!=nil{
	  log.Fatalln(err1)
  }
	_,err= insForm.Exec(developerid,"companyname","designation","location","startingdate","enddate","description","year","month","experience")
	if err!=nil{
	  log.Fatalln(err)
  }
		 http.Redirect(w,req,"/devaddprofile/personaldetails",307)
}

func postfreetrial(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	 defer db.Close()
	balance:=10000101101111101
	
	comid,_:=req.Cookie("Companyid")
	jobtitle,_:=req.Cookie("jobtitle")
  aboutcompany,_:=req.Cookie("aboutcompany")
  jobtype,_:=req.Cookie("jobtype")
	hires,_:=req.Cookie("hires")
	jobdescription,_:=req.Cookie("jobdescription")
	salary,_:=req.Cookie("salary")
	salarytypes,_:=req.Cookie("salarytypes")
	//stripeToken,_:=req.Cookie("stripeToken")
	skill1,_:=req.Cookie("skill1")
	skill2,_:=req.Cookie("skill2")
	skill3,_:=req.Cookie("skill3")
	otherskills,_:=req.Cookie("otherskills")
	tools,_:=req.Cookie("tools")
	minexp,_:=req.Cookie("minexp")
	maxexp,_:=req.Cookie("maxexp")

	
	rows,_:=db.Query("select companyname,location,website,companysize,industrytype,email,logo from companylogin WHERE idcompanylogin=?",comid.Value)
	for rows.Next(){
		err:=rows.Scan(&Companyname,&Companylocation,&Companywebsite,&Companysize,&Companytype,&Companyemail,&Companylogo)
			
		if err!=nil{
			log.Fatalln(err)
		}
	}	

	com:=[]string{jobtitle.Value,Companyname,Companylocation,Companywebsite,Companysize,Companytype,Companyemail,Companylogo,aboutcompany.Value,jobtype.Value,hires.Value,jobdescription.Value,salary.Value,salarytypes.Value,skill1.Value,skill2.Value,skill3.Value,otherskills.Value,tools.Value,minexp.Value,maxexp.Value}
insForm,err1 := db.Prepare("UPDATE companylogin SET balance= ? WHERE idcompanylogin=? ")
if err1!=nil{
			log.Fatalln(err1)
		 }
		 _,err:= insForm.Exec(0,comid.Value)
		if err!=nil{
						 log.Fatalln(err)
			 }
			 if balance==10000101101111101{
				 			 tpl.ExecuteTemplate(w,"stripePayment.tmpl",com)
			 }
	


}
func postconfirmation(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	 defer db.Close()
  comid,_:=req.Cookie("Companyid")
  jobtitle,_:=req.Cookie("jobtitle")
  aboutcompany,_:=req.Cookie("aboutcompany")
  jobtype,_:=req.Cookie("jobtype")
	hires,_:=req.Cookie("hires")
	jobdescription,_:=req.Cookie("jobdescription")
	salary,_:=req.Cookie("salary")
	salarytypes,_:=req.Cookie("salarytypes")
	//stripeToken,_:=req.Cookie("stripeToken")
	skill1,_:=req.Cookie("skill1")
	skill2,_:=req.Cookie("skill2")
	skill3,_:=req.Cookie("skill3")
	otherskills,_:=req.Cookie("otherskills")
	tools,_:=req.Cookie("tools")
	minexp,_:=req.Cookie("minexp")
	maxexp,_:=req.Cookie("maxexp")

	
	rows,_:=db.Query("select companyname,location,website,companysize,industrytype,email,logo,balance from companylogin WHERE idcompanylogin=?",comid.Value)
	for rows.Next(){
		err:=rows.Scan(&Companyname,&Companylocation,&Companywebsite,&Companysize,&Companytype,&Companyemail,&Companylogo,&Companybalance)
			
		if err!=nil{
			log.Fatalln(err)
		}
	}	

	com:=[]string{jobtitle.Value,Companyname,Companylocation,Companywebsite,Companysize,Companytype,Companyemail,Companylogo,aboutcompany.Value,jobtype.Value,hires.Value,jobdescription.Value,salary.Value,salarytypes.Value,skill1.Value,skill2.Value,skill3.Value,otherskills.Value,tools.Value,minexp.Value,maxexp.Value}

	if Companybalance!=0{
		tpl.ExecuteTemplate(w,"stripePayment.tmpl",com)
	}
	
}
func post(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	 defer db.Close()
	 comid,_:=req.Cookie("Companyid")
	 jobtitle,_:=req.Cookie("jobtitle")
	 aboutcompany,_:=req.Cookie("aboutcompany")
	 jobtype,_:=req.Cookie("jobtype")
	   hires,_:=req.Cookie("hires")
	   jobdescription,_:=req.Cookie("jobdescription")
	   salary,_:=req.Cookie("salary")
	   salarytypes,_:=req.Cookie("salarytypes")
	   //stripeToken,_:=req.Cookie("stripeToken")
	   skill1,_:=req.Cookie("skill1")
	   skill2,_:=req.Cookie("skill2")
	   skill3,_:=req.Cookie("skill3")
	   otherskills,_:=req.Cookie("otherskills")
	   tools,_:=req.Cookie("tools")
	   minexp,_:=req.Cookie("minexp")
	   maxexp,_:=req.Cookie("maxexp")
	currentTime := time.Now()
	status:=currentTime.Format("Jan-02-2006 Monday")
	fmt.Println(comid.Value,jobtitle.Value,Companyname,Companylocation,Companywebsite,Companysize,Companytype,Companyemail,Companylogo,aboutcompany.Value,jobtype.Value,hires.Value,jobdescription.Value,salary.Value,salarytypes.Value,skill1.Value,skill2.Value,skill3.Value,otherskills.Value,tools.Value,minexp.Value,maxexp.Value)
	rows,_:=db.Query("select companyname,location,website,companysize,industrytype,email,balance,stripeToken from companylogin WHERE idcompanylogin=?",comid.Value)
	for rows.Next(){
		err:=rows.Scan(&Companyname,&Companylocation,&Companywebsite,&Companysize,&Companytype,&Companyemail,&Companybalance,&Stripetoken)
			
		if err!=nil{
			log.Fatalln(err)
		}
	}	
balance:=Companybalance-1
insForm,err1 := db.Prepare("UPDATE companylogin SET balance= ? WHERE idcompanylogin=? ")
if err1!=nil{
	 log.Fatalln(err1)
	}
	_,err:= insForm.Exec(balance,comid.Value)
 if err!=nil{
					log.Fatalln(err)
		}

		insForm,err1 = db.Prepare("INSERT INTO company(idcompany,companyname,jobtitle,location,aboutcompany,size,jobtype,hires,jobdescription,industry,salary,salarytypes,website,companyemail,stripeToken,firstskill,secondskill,thirdskill,otherskill,toolsandtechnologies,minimumexp,maximumexp,postedon,status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err1!=nil{
		  log.Fatalln(err1)
	  }
		_,err= insForm.Exec(comid.Value,Companyname,jobtitle.Value,Companylocation,aboutcompany.Value,Companysize,jobtype.Value,hires.Value,jobdescription.Value,Companytype,salary.Value,salarytypes.Value,Companywebsite,Companyemail,Stripetoken,skill1.Value,skill2.Value,skill3.Value,otherskills.Value,tools.Value,minexp.Value,maxexp.Value,status,"open")
		if err!=nil{
		  log.Fatalln(err)
	  }
		http.Redirect(w,req,"/companyprofile",307)
}
func admin(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	 defer db.Close()
	err:= tpl.ExecuteTemplate(w,"admin.tmpl",nil)
			if err!=nil{
				  log.Fatalln(err)
			 }	
}	
func validateadmin(w http.ResponseWriter,req *http.Request ){
	db:=dbconnection.Connection()
	 defer db.Close()
	
	if req.Method == "GET" {
		var admin adminlogin

		Emailid:=req.FormValue("email")
		password:=req.FormValue("password")
		fmt.Println(Emailid,password)
		
		rows,_:=db.Query("select email,password from adminlogin WHERE email=? and password=?",Emailid,password)
			for rows.Next(){
				err:=rows.Scan(&admin.email,&admin.password)
				if err!=nil{
					log.Fatalln(err)
			   }	
			if admin.email==Emailid && admin.password==password{
				err:= tpl.ExecuteTemplate(w,"validateadmin.tmpl",nil)
				if err!=nil{
					log.Fatalln(err)
			   }	
			   fmt.Fprintln(w,"admin Logined")
			   
			}else{
				fmt.Fprintln(w,"email-id or  password is incorrect ")
			}
			}
		}
}
func validatelogin(w http.ResponseWriter,req *http.Request ){
	db:=dbconnection.Connection()
	 defer db.Close()
	
	if req.Method == "GET" {
		var company companylogin

		Emailid:=req.FormValue("email")
		Cname:=req.FormValue("cname")
		rows,_:=db.Query("select email,companyname from companylogin WHERE email=? and companyname=?",Emailid,Cname)
			for rows.Next(){
				err:=rows.Scan(&company.Email,&company.Companyname)
                if err!=nil{
					log.Fatalln(err)
			   }	
			if company.Email==Emailid && company.Companyname==Cname{
				err:= tpl.ExecuteTemplate(w,"validatelogin.tmpl",nil)
				if err!=nil{
					log.Fatalln(err)
			   }	
		}else{
				fmt.Fprintln(w,"email-id or  company name is incorrect ")
			}
			}
		}
}

func search(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	 defer db.Close()
	
	if req.Method == "GET" {
		var p person
		search:=req.FormValue("search")
		rows,_:=db.Query("select Company_name,Location,JobTitle from company WHERE Company_name=? or Location=? or JobTitle=? ",search,search,search)
			for rows.Next(){
				err:=rows.Scan(&p.Company,&p.Location,&p.JobTitle)
				if err!=nil{
					log.Fatalln(err)
			   }	
				//fmt.Fprintln(w,p.Company,p.Location,)
				tpl.ExecuteTemplate(w,"search.tmpl",&p)
			}
  
  	
}
}

func sendmail(w http.ResponseWriter,req *http.Request){

	/*auth := smtp.PlainAuth("", "krishchaituakula@gmail.com", "passwordd", "smtp.gmail.com")
	to := []string{"chaitanyaakula1910@gmail.com"}
	msg := []byte("To: chaitanyaakula1910@gmail.com\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "krishchaituakula@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}*/
	/*m := gomail.NewMessage()
	m.SetHeader("From", "info@gitty.app")
	m.SetHeader("To", "chaitanyaakula1910@gmail.com")
	//m.SetAddressHeader("Cc", "@gmail.com", "")
	m.SetHeader("Subject", "Gitty Apps")
	m.SetBody("text/plain", "Hello ")
	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "info@gitty.app", "Ancons@2020")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("mail send")
*/
from := mail.NewEmail("Gitty<info@gitty.app>", "info@gitty.app")
subject := "Recommended Jobs"
//replyto:=mail.NewEmail("reply","info@gitty.app")
to := mail.NewEmail("Developer", "krishchaituakula@gmail.com")
plainTextContent := "Recommended Jobs"
htmlContent := "<strong>Golang Developer ,Python Developer ,Java Developer</strong>"
message := mail.NewSingleEmail(from,subject, to, plainTextContent, htmlContent)
client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
response, err := client.Send(message)
if err != nil {
	log.Println(err)
} else {
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
}
}
func googlesignin(w http.ResponseWriter,req *http.Request){

	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
	
}
func profilepage(w http.ResponseWriter,req *http.Request){
	content, err := getUserInfo(req.FormValue("state"), req.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, req, "/signup", http.StatusTemporaryRedirect)
		return
	}
	arr := devprofilestruct{}
	if err = json.Unmarshal(content,&arr); err != nil {
		panic(err)
}
fmt.Println(arr)
	if err!=nil{
		log.Fatalln("error unmarshalling",err)
	}	
	Devemail:=arr.Email
	Devfirstname:=arr.GivenName
  Devlastname:=arr.FamilyName
	Devpicture:=arr.Picture
//	m,_ := time.ParseDuration("10m")
//	expire := time.Now().Add(m)  
	cookies.SetCookie(w,"devemail",Devemail)
	cookies.SetCookie(w,"devfirstname",Devfirstname)
  cookies.SetCookie(w,"devlastname",Devlastname)
  cookies.SetCookie(w,"devpicture",Devpicture)
  http.Redirect(w,req,"/newpassword",307)

}
func newpassword(w http.ResponseWriter,req *http.Request){

	devpicture,_:=req.Cookie("devpicture")
	devfirstname,_:=req.Cookie("devfirstname")
url:=devpicture.Value
response,e:=http.Get(url)
if e!=nil{
	log.Fatal(e)

}
defer response.Body.Close()
  fname:=devfirstname.Value
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	image:=fname+".jpg"
	path := filepath.Join(wd, "public", "pics",image)
fmt.Println(path)

	nf, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()
 _,err=io.Copy(nf,response.Body)
 if err!=nil{
	 log.Fatal(err)
 }
	tpl.ExecuteTemplate(w,"newpassword.tmpl",nil)

}
func confirmpassword(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	 defer db.Close()
	
	devmobile:=req.FormValue("mobile")
	devpassword:=req.FormValue("psw")
	devemailid,_:=req.Cookie("devemail")
	devfirstname,_:=req.Cookie("devfirstname")
	devlastname,_:=req.Cookie("devlastname")
	
	cookies.SetCookie(w,"devpassword",devpassword)
	cookies.SetCookie(w,"Emailid",devemailid.Value)
	cookies.SetCookie(w,"password",devpassword)
	insForm,err1 := db.Prepare("INSERT INTO developer(firstname,lastname,email,mobile,psw) VALUES(?,?,?,?,?)")
	if err1!=nil{
	log.Fatalln(err1)
}
	_,err:= insForm.Exec(devfirstname.Value,devlastname.Value,devemailid.Value,devmobile,devpassword)
	if err!=nil{
	log.Fatalln(err)
}


http.Redirect(w,req,"/verify1",307)
}
func editprofile2(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	 defer db.Close()
	
	devemailid,_:=req.Cookie("devemail")
	developerpassword,_:=req.Cookie("devpassword")
	rows,_:=db.Query("select iddeveloper,firstname,lastname,email,mobile from developer WHERE email=? and psw=?",	devemailid.Value,developerpassword.Value)
	for rows.Next(){
		err:=rows.Scan(&developerid,&developerFirstname,&developerLastname,&developerEmail,&developerMobile)
		if err!=nil{
			log.Fatalln(err)
	   }	
	}
	
cookies.SetCookie(w,"developerid",developerid)
rows1,_:=db.Query("select imagepath from profile WHERE iddeveloper=?",	developerid)
	for rows1.Next(){
		err:=rows1.Scan(&imagepath)
		if err!=nil{
			log.Fatalln(err)
	   }	
	}

chaitu:=[]string{imagepath,developerFirstname,developerLastname,developerEmail,developerMobile}
tpl.ExecuteTemplate(w,"editprofile2.tmpl",chaitu)
}
func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}


///////////////////////FILTER-END//////////////////
func jobresults(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	Cid:=req.FormValue("cid") 
	rows,_:=db.Query("select companyname,jobtitle,location,aboutcompany,jobdescription,salary,salarytypes,jobtype from company where companyuniqueid=? ",Cid)
	for rows.Next(){
		err:=rows.Scan(&Companyname,&Jobtitle,&Companylocation,&Aboutcompany,&Jobdesc,&Salary,&Salarytypes,&Jobtype)
		
		if err!=nil{
			log.Fatalln(err)
		}
	}
	searchresults:=[]string{Companyname,Jobtitle,Companylocation,Aboutcompany,Jobdesc,Salary,Salarytypes}
	tpl.ExecuteTemplate(w,"jobresults.tmpl",searchresults)
}
func freejobalert(w http.ResponseWriter,req *http.Request){
    tpl.ExecuteTemplate(w,"freejobalert.tmpl",nil)
}
func freejobalertsubscribed(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	username:=req.FormValue("username")
	 email:=req.FormValue("email")
	 mobile:=req.FormValue("mobile")
	 location:=req.FormValue("location")
	 skill1:=req.FormValue("skill1")
	 skill2:=req.FormValue("skill2")
	 skill3:=req.FormValue("skill3")
	 experience:=req.FormValue("experience")
	 insForm,err1 := db.Prepare("INSERT INTO freejobalert(username,email,mobile,location,skill1,skill2,skill3,experience) VALUES(?,?,?,?,?,?,?,?)")
	 if err1!=nil{
	   log.Fatalln(err1)
   }
	 _,err:= insForm.Exec(username,email,mobile,location,skill1,skill2,skill3,experience)
	 if err!=nil{
	   log.Fatalln(err)
   }


	
	tpl.ExecuteTemplate(w,"freejobalertsubscribed.tmpl",nil)

}
func about_GittyJobs(w http.ResponseWriter,req *http.Request){
	tpl.ExecuteTemplate(w,"aboutus.tmpl",nil)
}
func aboutGitty(w http.ResponseWriter,req *http.Request){
	tpl.ExecuteTemplate(w,"aboutGitty.tmpl",nil)
}
func contact_GittyJobs(w http.ResponseWriter,req *http.Request){
	tpl.ExecuteTemplate(w,"contactUs.tmpl",nil)
}
func contact_GittyJobs1(w http.ResponseWriter,req *http.Request){
    db:=dbconnection.Connection()
	defer db.Close()
	subject:=req.FormValue("subject")
	username:=req.FormValue("name")
	 email:=req.FormValue("email")
	 mobile:=req.FormValue("mobile")
	 location:=req.FormValue("location")
	 message:=req.FormValue("message")
	 insForm,err1 := db.Prepare("INSERT INTO contactus(subject,username,email,mobile,location,message) VALUES(?,?,?,?,?,?)")
	 if err1!=nil{
	   log.Fatalln(err1)
   }
	 _,err:= insForm.Exec(subject,username,email,mobile,location,message)
	 if err!=nil{
	   log.Fatalln(err)
   }
	tpl.ExecuteTemplate(w,"contactUs1.tmpl",nil)
}
func carrer(w http.ResponseWriter,req *http.Request){
	tpl.ExecuteTemplate(w,"carrers1.tmpl",nil)
}
func ourservices(w http.ResponseWriter,req *http.Request){
	tpl.ExecuteTemplate(w,"ourservices.tmpl",nil)
}

