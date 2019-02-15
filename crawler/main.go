package main

import (
	"file/learngo/crawler/engine"
	"file/learngo/crawler/persist"
	"file/learngo/crawler/scheduler"
	"file/learngo/crawler/skate/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParseCityList,
	//})
	itemChan, err := persist.ItemSaver("dating_show")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      98,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		//Url:    "http://www.zhenai.com/zhenghun",
		Url:    "http://www.chinashorttrack.com/jscs/",
		Parser: engine.NewFuncParser(parser.ParseContestList, "ParseCityList"),
	})

}
