package sendotp
import(
	"net/url"
	"fmt"
	"net/http"
	"io/ioutil"
)



func searchjobs(w http.ResponseWriter,req *http.Request){
	db:=dbconnection.Connection()
	defer db.Close()
	var details JobDetails1

	var jobdetails []JobDetails1=nil

	search,_:=req.Cookie("search")
	search_location,_:=req.Cookie("location")
	fulltime,_:=req.Cookie("fulltime")
	parttime,_:=req.Cookie("parttime")
	internship,_:=req.Cookie("internship")
	fmt.Println("testing...",search.Value,search_location.Value)
	


	if fulltime.Value==" "&& parttime.Value==" "&&internship.Value==" "{
		fmt.Println("tested")
		rows,_:=db.Query("select companyuniqueid,companyname,jobtitle,location,salary,salarytypes,jobtype,hires,jobdescription,industry,website,companyemail,firstskill,secondskill,thirdskill,otherskill,toolsandtechnologies,minimumexp,maximumexp,postedon from company where match(companyname,jobtitle) against(? IN  boolean mode) or location like ? ",search.Value,search_location.Value )
		
	for rows.Next(){
		
		err:=rows.Scan(&Companyid,&Companyname,&Jobtitle,&Companylocation,&Salary,&Salarytypes,&Jobtype,&Hires,&Jobdesc,&Industry,&Companywebsite,&Companyemail,&Companyskill1,&Companyskill2,&Companyskill3,&Companyotherskills,&Companytools,&Companyminexp,&Companymaxexp,&Companypostedon)
		details.Companyid=Companyid
		details.Companyname=Companyname
		details.Jobtitle=Jobtitle
		details.Location =Companylocation
		details.Salary=Salary
		details.Salarytypes=Salarytypes	
		details.Jobtype=Jobtype
     	details.Hires =Hires
	    details.Jobdesc=Jobdesc
	    details.Industry =Industry
     	details.Companywebsite=Companywebsite
     	details.Companyemail=Companyemail
    	details.Companyskill1=Companyskill1
    	details.Companyskill2 =Companyskill2
    	details.Companyskill3=Companyskill2
    	details.Companyotherskills=Companyotherskills
     	details.Companytools=Companytools
    	details.Companyminexp=Companyminexp
    	details.Companymaxexp=Companymaxexp
    	details.Companypostedon=Companypostedon
		details.Count_fulltime=" "
		details.Count_parttime=" "
		details.Count_internship=" "
		jobdetails=append(jobdetails,details)
		if err!=nil{
			log.Fatalln(err)
		}
	
	}

}
	if fulltime.Value=="Full-time"{
		fmt.Println(" fulltime tested")
		  
	    rows,_:=db.Query("select companyuniqueid,companyname,jobtitle,location,salary,salarytypes,jobtype,hires,jobdescription,industry,website,companyemail,firstskill,secondskill,thirdskill,otherskill,toolsandtechnologies,minimumexp,maximumexp,postedon from company where (match(companyname,jobtitle) against(? IN  boolean mode) or  location like ?) and jobtype like ? ",search.Value,search_location.Value,fulltime.Value)
		for rows.Next(){
		
			err:=rows.Scan(&Companyid,&Companyname,&Jobtitle,&Companylocation,&Salary,&Salarytypes,&Jobtype,&Hires,&Jobdesc,&Industry,&Companywebsite,&Companyemail,&Companyskill1,&Companyskill2,&Companyskill3,&Companyotherskills,&Companytools,&Companyminexp,&Companymaxexp,&Companypostedon)
			details.Companyid=Companyid
			details.Companyname=Companyname
			details.Jobtitle=Jobtitle
			details.Location =Companylocation
			details.Salary=Salary
			details.Salarytypes=Salarytypes	
			details.Jobtype=Jobtype
			 details.Hires =Hires
			details.Jobdesc=Jobdesc
			details.Industry =Industry
			 details.Companywebsite=Companywebsite
			 details.Companyemail=Companyemail
			details.Companyskill1=Companyskill1
			details.Companyskill2 =Companyskill2
			details.Companyskill3=Companyskill2
			details.Companyotherskills=Companyotherskills
			 details.Companytools=Companytools
			details.Companyminexp=Companyminexp
			details.Companymaxexp=Companymaxexp
			details.Companypostedon=Companypostedon
			details.Count_fulltime=" "
			details.Count_parttime=" "
			details.Count_internship=" "
			jobdetails=append(jobdetails,details)
			if err!=nil{
				log.Fatalln(err)
			}
		
		}
	}
	if parttime.Value=="Part-time"{
		 fmt.Println("parttime tested")
		  
	    rows,_:=db.Query("select companyuniqueid,companyname,jobtitle,location,salary,salarytypes,jobtype,hires,jobdescription,industry,website,companyemail,firstskill,secondskill,thirdskill,otherskill,toolsandtechnologies,minimumexp,maximumexp,postedon from company where (match(companyname,jobtitle) against(? IN  boolean mode) or  location like ?) and jobtype like ? ",search.Value,search_location.Value,parttime.Value)
		for rows.Next(){
		
			err:=rows.Scan(&Companyid,&Companyname,&Jobtitle,&Companylocation,&Salary,&Salarytypes,&Jobtype,&Hires,&Jobdesc,&Industry,&Companywebsite,&Companyemail,&Companyskill1,&Companyskill2,&Companyskill3,&Companyotherskills,&Companytools,&Companyminexp,&Companymaxexp,&Companypostedon)
			details.Companyid=Companyid
			details.Companyname=Companyname
			details.Jobtitle=Jobtitle
			details.Location =Companylocation
			details.Salary=Salary
			details.Salarytypes=Salarytypes	
			details.Jobtype=Jobtype
			 details.Hires =Hires
			details.Jobdesc=Jobdesc
			details.Industry =Industry
			 details.Companywebsite=Companywebsite
			 details.Companyemail=Companyemail
			details.Companyskill1=Companyskill1
			details.Companyskill2 =Companyskill2
			details.Companyskill3=Companyskill2
			details.Companyotherskills=Companyotherskills
			 details.Companytools=Companytools
			details.Companyminexp=Companyminexp
			details.Companymaxexp=Companymaxexp
			details.Companypostedon=Companypostedon
			details.Count_fulltime=" "
			details.Count_parttime=" "
			details.Count_internship=" "
			jobdetails=append(jobdetails,details)
			if err!=nil{
				log.Fatalln(err)
			}
		
		}
	}
	if internship.Value=="Internship"{
		 
		  fmt.Println("internship tested")
	    rows,_:=db.Query("select companyuniqueid,companyname,jobtitle,location,salary,salarytypes,jobtype,hires,jobdescription,industry,website,companyemail,firstskill,secondskill,thirdskill,otherskill,toolsandtechnologies,minimumexp,maximumexp,postedon from company where (companyname like ? or jobtitle like ? or  location like ?) and jobtype like ? ",search.Value,search.Value,search_location.Value,internship.Value)
		for rows.Next(){
		
			err:=rows.Scan(&Companyid,&Companyname,&Jobtitle,&Companylocation,&Salary,&Salarytypes,&Jobtype,&Hires,&Jobdesc,&Industry,&Companywebsite,&Companyemail,&Companyskill1,&Companyskill2,&Companyskill3,&Companyotherskills,&Companytools,&Companyminexp,&Companymaxexp,&Companypostedon)
			details.Companyid=Companyid
			details.Companyname=Companyname
			details.Jobtitle=Jobtitle
			details.Location =Companylocation
			details.Salary=Salary
			details.Salarytypes=Salarytypes	
			details.Jobtype=Jobtype
			 details.Hires =Hires
			details.Jobdesc=Jobdesc
			details.Industry =Industry
			 details.Companywebsite=Companywebsite
			 details.Companyemail=Companyemail
			details.Companyskill1=Companyskill1
			details.Companyskill2 =Companyskill2
			details.Companyskill3=Companyskill2
			details.Companyotherskills=Companyotherskills
			 details.Companytools=Companytools
			details.Companyminexp=Companyminexp
			details.Companymaxexp=Companymaxexp
			details.Companypostedon=Companypostedon
			details.Count_fulltime=" "
			details.Count_parttime=" "
			details.Count_internship=" "
			jobdetails=append(jobdetails,details)
			if err!=nil{
				log.Fatalln(err)
			}
		
		}
	}
	rows1,_:=db.Query("select (select count(*) from company where jobtype like ?)as count_fulltime,( select count(*)  from company where jobtype like ? )as count_parttime,( select count(*)  from company where jobtype like ?)as count_internship","Full-time","Part-time","Internship")
	for rows1.Next(){
				err:=rows1.Scan(&Count_fulltime,&Count_parttime,&Count_internship)
				details.Count_fulltime=Count_fulltime
				details.Count_parttime=Count_parttime
				details.Count_internship=Count_internship
				if err!=nil{
					log.Fatalln(err)
				}
		}
		jobdetails=append(jobdetails,details)
	
fmt.Println(jobdetails)
	tpl.ExecuteTemplate(w,"search.tmpl",jobdetails)
}
///////////////////////FILTERS//////////////////////

func filter_fulltime(w http.ResponseWriter,req *http.Request){
	
	cookies.SetCookie(w,"fulltime","Full-time")
	cookies.SetCookie(w,"parttime"," ")
	cookies.SetCookie(w,"internship"," ")
	http.Redirect(w,req,"/searchjobs",307)
}
func filter_parttime(w http.ResponseWriter,req *http.Request){
	cookies.SetCookie(w,"parttime","Part-time")
	cookies.SetCookie(w,"fulltime"," ")
	cookies.SetCookie(w,"internship"," ")
	http.Redirect(w,req,"/searchjobs",307)
}
func filter_internship(w http.ResponseWriter,req *http.Request){
	cookies.SetCookie(w,"internship","Internship")
	cookies.SetCookie(w,"fulltime"," ")
	cookies.SetCookie(w,"parttime"," ")
	http.Redirect(w,req,"/searchjobs",307)
}






///////////////////////FILTER-END//////////////////