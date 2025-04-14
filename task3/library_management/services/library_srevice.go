package services
import (
	"fmt"
	"library_management/models"
)

 type Library struct{
	books map[int]models.Book
	members map[int]models.Member

}
type  LibraryManager interface{
	AddBook(models.Book)
	RemoveBook(bookid int)
	// it will allow users to  borrow the book
	BorrowBook(bookid int, memberid int)
	// it will allow users   to return  all the book
	ReturnBook(bookid int,memberid int)
	// list  the unborrowed book
	ListAvaliableBook()[] models.Book
	//list the borrowed book
	ListBorrowedBook(memberID int) []models.Book
	AddMember(models.Member)
}
// here  we will implement the  method , the  is going 
func NewLibrary() *Library{
	return &Library{
		books: make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}
func (l *Library) AddBook(book models.Book){
	l.books[book.ID]=book
	fmt.Printf("%s", "book added successfully")

}
func  (l *Library) RemoveBook(bookid int){
	   delete(l.books,bookid)
}

func  (l *Library)  BorrowBook(bookid  int, memeberid  int){
	newbook := l.books[bookid]
	if newbook.Status=="Available"{
		var newmemeber models.Member;
		if member, exists := l.members[memeberid]; exists{
		  newmemeber = member
	}else{
		fmt.Println("member is not registered");
		return 
		
	}
	    newbook.Status="Borrowed"
		newmemeber.BorrowedBooks = append(newmemeber.BorrowedBooks, newbook)
		l.members[memeberid]=newmemeber
		l.books[bookid]=newbook
	}
    
}
func (l *Library)  ReturnBook(bookid int, memeberid int){
	 newbook:= l.members[memeberid].BorrowedBooks[bookid]
	 newmemeber:= l.members[memeberid]
	 var newBookOfusers []models.Book

		// newmemeber.BorrowedBooks = remove(newmemeber.BorrowedBooks, newbook)
		 
	 for _,val := range newmemeber.BorrowedBooks{
		if val != newbook{
			newBookOfusers=append(newBookOfusers,val)
           
		}
	 }
	  previousbook:= l.books[bookid]
	  previousbook.Status="Available"
	  l.books[bookid]=previousbook
      newmemeber.BorrowedBooks=newBookOfusers
	  l.members[memeberid]=newmemeber

	}
func (l*Library) ListAvaliableBook() []models.Book{
	var newavliablebooks []models.Book
	for _, book := range l.books{
		if  book.Status=="Available"{
			newavliablebooks = append(newavliablebooks, book)

		}

	}
    return newavliablebooks
}
func (l*Library) ListBorrowedBook(memberid int) []models.Book{
	var newborrowedbooks []models.Book
	if member,exists :=l.members[memberid];exists{
		newborrowedbooks = member.BorrowedBooks
}else{
		fmt.Println("user does not exist");
	}
	return newborrowedbooks

}
func (l*Library) AddMember(memeber models.Member){
	l.members[memeber.ID]=memeber
	fmt.Printf("%s", "member added successfully")


}