package main

import (
	"testing"
	"file/learngo/crawler_distributed/rpcsupport"
	"file/learngo/crawler_distributed/worker"
	"time"
	"file/learngo/crawler_distributed/config"
	"fmt"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/96307513",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "男人",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	}else {
		fmt.Println(result)
	}
}
