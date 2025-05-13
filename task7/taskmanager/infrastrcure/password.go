package infrastrcure
import(
	"golang.org/x/crypto/bcrypt"
	"taskmanager/domain"
	"log"

)
type   User  struct{
	user  domain.User
}
func LoginChekcer(newuser domain.User, password string )  string{
	usersHashedpassword := newuser.Password
	bcrypterr:=bcrypt.CompareHashAndPassword([]byte(usersHashedpassword),[]byte(password))
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