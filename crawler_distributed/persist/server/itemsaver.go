package main

import (
	"file/learngo/crawler_distributed/rpcsupport"
	"file/learngo/crawler_distributed/persist"
	"github.com/olivere/elastic"
	"file/learngo/crawler_distributed/config"
	"log"
	"fmt"
	"flag"
)
var port = flag.Int("port",0,"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0{
		fmt.Println("must specify a prot")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d",*port),config.ElasticIndex))
}
func serveRpc(host,index string) error{
	client,err:=elastic.NewClient(elastic.SetURL("http://192.168.99.100:9200/"),elastic.SetSniff(false))
	if err != nil{
		return err
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
		Client:client,
		Index:index,
	})
}
