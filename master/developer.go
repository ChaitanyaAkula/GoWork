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