package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {
	resp,err := http.Get("https://m.bilibili.com/space/104207471")
	if err!= nil{
		panic(err)
	}
	defer resp.Body.Close()
	s,err := httputil.DumpResponse(resp,true)
	if err !=nil{
		panic(err)
	}
	fmt.Printf("%s", s)
}
