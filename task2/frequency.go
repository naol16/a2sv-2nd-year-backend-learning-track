package main
import ("fmt"
"bufio"
"os"
"strings"
"unicode"
)

func repitionCounter(thestring string) map[string]int {
    counter := make(map[string]int)
    for _, char := range thestring {
		key := string(char)
        counter[key]++
    }
    return counter
}

func  palindromechecker(newstring string) bool{
	newstring=toLowerCase(newstring)
	newstring=removeSpaces(newstring)
	newstring=removePunctuation(newstring)
	var index=len(newstring)-1
	for i:=0;i<len(newstring);i++{
		if newstring[i]!=newstring[index-i]{
			return false
		}

	}
	return true
}
func toLowerCase(s string) string {
    return strings.ToLower(s)
}
func removeSpaces(s string) string {
    var b strings.Builder
    for _, r := range s {
        if r != ' ' {
            b.WriteRune(r)
        }
    }
    return b.String()
}
func removePunctuation(s string) string {
    var b strings.Builder
    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            b.WriteRune(r)
        }
    }
    return b.String()
}
func main(){
	var letter string
	fmt.Println("enter the letter")
	reader := bufio.NewReader(os.Stdin) 
	letter, _ = reader.ReadString('\n')
	fmt.Println(
	palindromechecker(letter),repitionCounter(letter))

}