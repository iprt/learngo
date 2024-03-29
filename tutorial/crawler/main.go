package main

import (
	"iproute.org/learngo/tutorial/crawler/config"
	"iproute.org/learngo/tutorial/crawler/engine"
	"iproute.org/learngo/tutorial/crawler/persist"
	"iproute.org/learngo/tutorial/crawler/scheduler"
	"iproute.org/learngo/tutorial/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
