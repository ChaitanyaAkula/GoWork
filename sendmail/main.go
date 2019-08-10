package main
import(
	"gopkg.in/gomail.v2"
	"fmt"
)
func main(){
	m := gomail.NewMessage()
m.SetHeader("From", "from@gmail.com")
m.SetHeader("To", "To@gmail.com")
//m.SetAddressHeader("Cc", "@gmail.com", "")
m.SetHeader("Subject", "Gitty Apps")
m.SetBody("text/plain", "Hello ")
d := gomail.NewPlainDialer("smtp.gmail.com", 587, "from@gmail.com", "password")
if err := d.DialAndSend(m); err != nil {
    panic(err)
}
fmt.Println("mail send")
}