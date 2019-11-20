package reverse
import(
"strconv"
"assignment/loop"
)
type Runes []rune
 
func (str Runes) ReverseString() (revStr Runes) {
    l := len(str); revStr = make(Runes, l)    
    for i := 0; i <= l/2; i++ { 
        revStr[i], revStr[l-1-i] = str[l-1-i], str[i] 
    }
    return revStr
}
 
func (str Runes) String() string {
    return string(str)
}
 
func Formreverse()string{
	for i := 0; i <= 9; i++ {
	return strconv.Itoa(loop.FibonacciUsingLoops(i)) + " "
		
    }
    return("")

}