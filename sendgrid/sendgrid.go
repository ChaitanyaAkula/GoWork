package resumeupload
import(
	"fmt"
	"net/http"
	
)



func sendgridt(w http.ResponseWriter, req *http.Request ){
	
	 
	// sendotp.Send(mob1,Msg)
	from := mail.NewEmail("Gitty", "from@gmail.com")
	// replyto:=mail.NewEmail("reply","from@gmail.com")
  subject := "OTP Verification for GittyJobs"
  to := mail.NewEmail("Company","To@gmail.com")
  plainTextContent := "Enter the below OTP in verification page  "
  htmlContent := ` <p> OTP:</p> `+str+` ` 
  
  
  message := mail.NewSingleEmail(from,subject, to, plainTextContent, htmlContent)
  client := sendgrid.NewSendClient(SENDGRID_KEY)
  response, err := client.Send(message)
  if err != nil {
	  log.Println(err)
  } else {
	  fmt.Println(response.StatusCode)
	  fmt.Println(response.Body)
	  fmt.Println(response.Headers)
  }