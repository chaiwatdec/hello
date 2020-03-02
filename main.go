package main
import s "strings"
import "fmt"

var p = fmt.Println
func main() {
	p("Contains:	", s.Contains("test", "es"))
	p("Count:		", s.Count("test", "t"))	//count t
	p("HasPrefix:	", s.HasPrefix("test", "te"))	//start with te?
	p("HasSuffix:	", s.HasSuffix("test", "st"))	//end with st?
	p("Index:		", s.Index("test", "t"))	//find position
	p("Join:		", s.Join([]string{"a", "b"}, "-")) //join a a with -
	p("Repeat:		", s.Repeat("a", 5)) //repeat 1 5 times
	p("Replace:	", s.Replace("foo", "o", "0", -1)) 
	p("Replace:	", s.Replace("fooo", "o", "0", 2))	//replace o with 0 2 time
	p("Split:		", s.Split("a-b-c-d-e", "-")) //split with -
	p("ToUpper:	", s.ToUpper("test"))
	p("Len:		", len("hello")) //count string
}

