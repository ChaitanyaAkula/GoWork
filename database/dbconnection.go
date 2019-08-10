package dbconnection
import(
	"log"
	_"github.com/go-sql-driver/mysql"
	
	"database/sql"
)
var db *sql.DB
func Connection()*sql.DB{

	Db,err := sql.Open("mysql","root:root@tcp(localhost:3306)/gitty")
	
	if err!=nil{
		log.Fatalln(err)
	}
	return Db
}
