package main

import (
	"file/learngo/crawler_distributed/rpcsupport"
	"fmt"
	"file/learngo/crawler_distributed/worker"
	"log"
	"flag"
)

var port = flag.Int("port",0,"the port for me to listen on")

func main(){
	flag.Parse()
	if *port == 0{
		fmt.Println("must specify a prot")
		return
	}
	log.Fatal(
		rpcsupport.ServeRpc(fmt.Sprintf(":%d",*port),worker.CrawlService{}))
}
