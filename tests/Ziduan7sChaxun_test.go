package tests

import (
	"log"
	"mhsyquanxian/ceshiyong"
	"mhsyquanxian/ziduan7skus"
	"testing"
)

func TestChaxunZiduan7s(t *testing.T) {
	ret := ziduan7skus.ChaxunZiduan7s(ceshiyong.Ziduan7schaxun())
	log.Println(ret)
}
