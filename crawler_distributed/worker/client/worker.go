package client

import (
	"file/learngo/crawler/engine"
	"file/learngo/crawler_distributed/config"
	"file/learngo/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor) {


	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult

		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{},nil
		}
		return worker.DeserializeResult(sResult), nil
	}
}
