package main

import (
	"fmt"

	"log"

	"flag"

	"iproute.org/learngo/tutorial/crawler/fetcher"
	"iproute.org/learngo/tutorial/crawler_distributed/rpcsupport"
	"iproute.org/learngo/tutorial/crawler_distributed/worker"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	fetcher.SetVerboseLogging()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
