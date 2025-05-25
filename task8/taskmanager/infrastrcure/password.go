package infrastrcure
import(
	"golang.org/x/crypto/bcrypt"
	"log"

)
func LoginChekcer(newuserpassword string, password string )  string{
	bcrypterr:=bcrypt.CompareHashAndPassword([]byte(newuserpassword),[]byte(password))
	if bcrypterr!=nil{
		return "either your password or username does not much"

}
 return ""
}
func  Hasher (password  string) []byte{
	hashedpassword,err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err!=nil{
		log.Fatal(err)
	}
	
	return hashedpassword


}