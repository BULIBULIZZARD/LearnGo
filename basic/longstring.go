package main

import (
	"fmt"
)

func myLongstring(s string)  {

	long := ""
	string :=""
	start  := 0
	for i:=0 ;i<len(s);i++{
		for j:=start;j<i;j++{
			if(s[j]==s[i]){
				start = i
			}
		}
		string =s[start:i+1]
		if len(string)>len(long) {
			long = string
		}
	}
	fmt.Println(long)
}
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make (map[rune] int)
	start := 0
	maxLength := 0
	for i,ch:= range []rune(s){

		if lastI,ok := lastOccurred[ch]; ok&&lastI >= start {
			start = lastI +1
		}
		if i - start + 1 >maxLength{
			maxLength = i - start +1
		}
		lastOccurred[ch] =i
	}
	return maxLength
}
func main() {
	//myLongstring("qweqweasdfghjk")
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	myLongstring("abcabcbb")
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	myLongstring("bbbbb")
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	myLongstring("pwwkew")
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	myLongstring("")
	fmt.Println(lengthOfNonRepeatingSubStr("asdfghj"))
	myLongstring("asdfghj")
	fmt.Println(lengthOfNonRepeatingSubStr("你是真的鬼畜"))
	myLongstring("你是真的鬼畜")


}
