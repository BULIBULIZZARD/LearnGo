package main

import (
	"fmt"
)

func tryRecover(){
	defer func() {
		r := recover()
		if err , ok := r.(error); ok{
			fmt.Println("Error occurred ",err)
		}else{
			panic(fmt.Sprintf("what ? %v",r))
		}
	}()
	panic(123)
}
func main() {
	tryRecover()
}
