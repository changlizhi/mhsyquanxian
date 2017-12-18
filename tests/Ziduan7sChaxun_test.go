package tests

import (
	"testing"
	"mhsyquanxian/ziduan7skus"
	"mhsyquanxian/ceshiyong"
	"log"
)

func TestChaxunZiduan7s(t *testing.T) {
	ret := ziduan7skus.ChaxunZiduan7s(ceshiyong.Ziduan7schaxun())
	log.Println(ret)
}
