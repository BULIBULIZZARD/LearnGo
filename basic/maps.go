package main

import "fmt"

func main() {

	m := map[string]string{
		"name":"ccmouse",
		"course":"golang",
		"site":"imooc",
		"quality":"notbad",
	}
	fmt.Println(m)
	for k,v := range m {
		fmt.Println(k,v)
	}
}
