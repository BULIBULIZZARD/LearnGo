package main

import (
	"regexp"
	"fmt"
)

const text  = "my email is 978082243@qq.com"
func main(){
	re := regexp.MustCompile(`[a-zA-z0-9_]+@[a-zA-z0-9.]+\.[a-zA-z0-9]+`)
	match := re.FindString(text)
	fmt.Println(match)
}
