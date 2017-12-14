package quanju

import (
	"log"
	"mhsyquanxian/gongjus"
)

var flakeid *gongjus.IdWorker

func Flakeid() *gongjus.IdWorker {
	if flakeid == nil {
		flakeid, err := gongjus.NewIdWorker()
		if err != nil {
			log.Println("flakeid shengcheng shibai", err)
		}
		log.Println("Flakeid-----", flakeid)
		return flakeid
	}
	return flakeid
}
