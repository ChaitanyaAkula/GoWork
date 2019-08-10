package devprofileupload
import(
	"fmt"
	"net/http"
	"os"
	"io"
	"path/filepath"
	"crypto/sha1"
	"strings"
    "mime/multipart"
)



func SetProfile(w http.ResponseWriter,mf multipart.File,fh *multipart.FileHeader)string{

	defer mf.Close()
	ext := strings.Split(fh.Filename, ".")[1]
	h := sha1.New()
	io.Copy(h, mf)
	fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "public", "developerprofile", fname)
	nf, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()
	mf.Seek(0, 0)
	io.Copy(nf, mf)
	fmt.Println("chaitanya")
	return fname
}