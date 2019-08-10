package cookies
import "net/http"


func SetCookie(w http.ResponseWriter,name,value string){
        http.SetCookie(w,&http.Cookie{
		Name:name,
		Value:value,
	//	Expires:expire,
		HttpOnly: true,
	})

}