package googlesignin
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

func GoogleSignin(){
	func init(){


		tpl = template.Must(template.New("").ParseGlob("templates/*.tmpl"))
		googleOauthConfig = &oauth2.Config{
			RedirectURL:  "http://localhost:8080/profile",
		   ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		   ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email","profile"},
			Endpoint:     google.Endpoint,
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
}