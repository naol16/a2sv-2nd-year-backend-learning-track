package services
import(
	"library_management/models"
)
type  LibraryManager interface{
	AddBook()
	RemoveBook()
	BorrowBook()
	ReturnBook()
	ListAvaliableBook()
	ListBorrowedBook()

}
// here  we will implement the  method , the  is going 

func (newbook Book)AddBook(){

}
func  (bookid int) RemoveBook(){

}