package sendotp
import(
	"net/url"
	"fmt"
	"net/http"
	"io/ioutil"
)



type Details struct {
	Username string
	Userid   string
	Handle   string
	Msg      int64
	From     string
	To       string
	Customid string
	Price    int
	Mccmnc   int
	Credit   int
}

func Send(mob1 string,Msg int64){
	var detail Details
  
	fmt.Println(mob1)
	detail.Username="Chaitanyaakula10"
	detail.Userid="User_id"
	detail.Handle="Handle"
	detail.Msg=Msg
	detail.From="Name"
	detail.To=mob1
	detail.Customid=""
	detail.Price=0
	detail.Mccmnc =0
	detail.Credit=0
	
	
	
	url:=buildURL(detail)
	fmt.Println(url)
	response,err:=http.Get(url)
	if err!=nil{
		fmt.Println("error",err)
		return
	}
	defer response.Body.Close()
	body,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Println("Error",err)
		return
	}
	
	fmt.Println(body)
	return 
}

func buildURL(det Details) string {
	
	u := &url.URL{
		Scheme:   "https",
		Host:     "api.budgetsms.net",
		Path:     "/sendsms",

		RawQuery: fmt.Sprintf("username=%s&userid=%s&handle=%s&msg=Your+OTP+for+GittyAppss+is+%d&from=%s&to=%s&customid=%s&price=%d&mccmnc=%d&credit=%d", det.Username, det.Userid, det.Handle,det.Msg, det.From, det.To, det.Customid, det.Price, det.Mccmnc, det.Credit),
	}

	return u.String()
}