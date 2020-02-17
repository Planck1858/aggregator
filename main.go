package main

import (
	"github.com/Planck1858/aggregator/news"
	"github.com/Planck1858/aggregator/router"
)

func main() {
	r := router.New()
	a := news.New()

	go a.Serve()
	r.Run()
}
