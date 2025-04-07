package main
import ("fmt"
"bufio"
"os"
"strconv"
"strings"
)



type Subject struct{
	subject string
	grade   int
}
type student struct{
	Name string
	subjects  []Subject

}


func averagecalutaor(subjectOfUser []Subject) int{
	var totalsum int
	for i:=0; i<len(subjectOfUser);i++{
		totalsum+=subjectOfUser[i].grade

	};
	average:= totalsum/len(subjectOfUser)
	return average

 }

func main(){
var name string
var newUser student
numberOfSubjects:=0;
reader := bufio.NewReader(os.Stdin) // Use bufio to read full names
fmt.Println("Enter your full name:")
name, _ = reader.ReadString('\n')
fmt.Scanln(&name)  // Read full line (including spaces)
fmt.Println("Enter number of subjects:")
fmt.Scanln(&numberOfSubjects) 
newUser =student{Name:name,subjects:[]Subject{}}
for i:=0; i<numberOfSubjects;i++{
	var subject string
	var grade int
	var newSubject Subject
    reader := bufio.NewReader(os.Stdin) // Use bufio to read full name                           // Read full line (including spaces)
    fmt.Println("Enter name of the subject")
    fmt.Scanln(&subject)
	   // Use bufio to read full names
    fmt.Println("Enter the grade")
    input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
    grade, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("Invalid input. Please enter a valid number.")
		return
} // Read full line (including spaces)
if   grade>100 || grade<0{
		fmt.Println("you should have the correct grade")
		return

}
newSubject=Subject{subject:subject,grade:grade}
newUser.subjects=append(newUser.subjects,newSubject)

fmt.Println("student name",newUser.Name)
for i:=0; i<len(newUser.subjects) ;i++{
	fmt.Println("subject",newUser.subjects[i].subject,"grade",newUser.subjects[i].grade)
}

fmt.Println("average grade",averagecalutaor(newUser.subjects))


}
}