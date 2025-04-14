package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

//  here we will implement the logic based on the number
// 1--->

func  command(){
	fmt.Println("\nLibrary Management System")
	fmt.Println("1. Add a new book")
	fmt.Println("2. Remove a book")
	fmt.Println("3. Borrow a book")
	fmt.Println("4. Return a book")
	fmt.Println("5. List all available books")
	fmt.Println("6. List all borrowed books by member")
	fmt.Println("7. Add a new member")
	fmt.Println("8. Exit")
	fmt.Print("Select an option (1-7): ")

}
type LibraryController struct {
    service services.LibraryManager
}
func NewController(service services.LibraryManager) *LibraryController {
    return &LibraryController{service: service}
}
func (lc *LibraryController)addBook(reader bufio.Reader){
	fmt.Printf("%s","enter your bookid")
	bookid,_ := reader.ReadString('\n')
	bookid = strings.TrimSpace(bookid)
	newbookid, err := strconv.Atoi(bookid)
	if err!=nil{
		fmt.Printf("%s","enter  the correct id")
		return

	}
	fmt.Printf("%s","enter you title")
	title, err:=reader.ReadString('\n')
	for  err!=nil{
		fmt.Printf("%s","enter the correct title")
		return
	}
	fmt.Printf("%s","enter the author")
	author, err:=reader.ReadString('\n')
	for err!=nil{
		fmt.Printf("%s","enter the author if not err")
		return
	}
    newbook := models.Book{
  	ID : newbookid,
	Title :title,
	Author :author,
	Status :"Available"}
 lc.service.AddBook(newbook)

}

func(lc*LibraryController) removeBook(reader bufio.Reader){   
   
	fmt.Printf("%s","enter your bookid")
	bookid,_ := reader.ReadString('\n')
	bookid = strings.TrimSpace(bookid)
	newbookid, err := strconv.Atoi(bookid)
	if err!=nil{
		fmt.Printf("%s","enter  the correct id")
		return
	}
	lc.service.RemoveBook(newbookid)

}
func(lc*LibraryController) borrowBook(reader bufio.Reader){
	fmt.Printf("%s","enter your bookid")
	bookid,_ := reader.ReadString('\n')
	bookid = strings.TrimSpace(bookid)
	newbookid, err := strconv.Atoi(bookid)
	if err!=nil{
		fmt.Printf("%s","enter  the correct id")
		return
	}
	fmt.Printf("%s","enter your memberid")
	memberid,_ := reader.ReadString('\n')
	memberid = strings.TrimSpace(memberid)
	newmemberid, err:= strconv.Atoi(memberid)
	if err!=nil{
		fmt.Printf("%s","enter  the correct id")
		return
	}
	lc.service.BorrowBook(newbookid,newmemberid)

}
func(lc*LibraryController) returnBook(reader bufio.Reader){
	fmt.Printf("%s","enter your bookid")
	bookid,_ := reader.ReadString('\n')
	bookid = strings.TrimSpace(bookid)
	newbookid, err := strconv.Atoi(bookid)
	if err!=nil{
		fmt.Printf("%s","enter  the correct id")
		return
	}
	fmt.Printf("%s","enter your memberid")
	memberid,_ := reader.ReadString('\n')
	memberid = strings.TrimSpace(memberid)
	newmemberid, err:= strconv.Atoi(memberid)
	if err!=nil{
		fmt.Printf("%s","enter  the correct id")
		return
	}
	lc.service.ReturnBook(newbookid,newmemberid)


}
func(lc*LibraryController) listAvailableBooks(){
   	availableBooks := lc.service.ListAvaliableBook()
	fmt.Println(availableBooks)
	if len(availableBooks) == 0 {
		fmt.Println("No available books.")
		return
	}
	fmt.Println("Available Books:")
	for _, book := range availableBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}	

}
func (lc*LibraryController)AddMember(reader bufio.Reader){
	fmt.Printf("%s","enter your member id")
	memberid,_ := reader.ReadString('\n')
	memberid = strings.TrimSpace(memberid)
	userid, err := strconv.Atoi(memberid)
	if err!=nil{
		fmt.Printf("%s","enter  the correct member id")
		return

	}
	fmt.Printf("%s","enter you name")
	name, err:=reader.ReadString('\n')
	for  err!=nil{
		fmt.Printf("%s","enter the correct name")
		return
	}
	var newuser models.Member
	newuser = models.Member{
		ID:userid,
		Name:name,
		BorrowedBooks:[]models.Book{},
	}
	lc.service.AddMember(newuser)

}
func(lc*LibraryController)listBorrowedBooksByMember(reader bufio.Reader){
	fmt.Printf("%s","enter your memberid")
	memberid,_ := reader.ReadString('\n')
	memberid = strings.TrimSpace(memberid)
	newmemberid, err:= strconv.Atoi(memberid)
	if err!=nil{
		fmt.Printf("%s","enter  the correct id")
		return
	}
	borrowedBooks := lc.service.ListBorrowedBook(newmemberid)
	fmt.Println(borrowedBooks)
	if len(borrowedBooks) == 0 {
		fmt.Println("No borrowed books.")
		return
	}
	fmt.Println("Borrowed Books:")
	for _, book := range borrowedBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}


}

func (lc*LibraryController) Run(){
reader := bufio.NewReader(os.Stdin)
for {
    command() 
	var choice string;
	choice, _=reader.ReadString('\n')
	choice= strings.TrimSpace(choice)
	switch choice {
	case "1":
		lc.addBook(*reader)
	case "2":
		lc.removeBook(*reader)
	case "3":
		lc.borrowBook(*reader)
	case "4":
		lc.returnBook(*reader)
	case "5":
		lc.listAvailableBooks()
	case "6":
		lc.listBorrowedBooksByMember(*reader)
	case "7":
		lc.AddMember(*reader)
	case "8":
		fmt.Println("Exiting the system...")
		os.Exit(0)
	default:
		fmt.Println("Invalid option. Please try again.")

	}
}


}
