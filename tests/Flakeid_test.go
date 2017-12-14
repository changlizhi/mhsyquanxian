package tests

import (
	"log"
	"mhsyquanxian/quanju"
	"testing"
)

func TestFlakeid(t *testing.T) {
	id, _ := quanju.Flakeid().NextId()
	log.Println("id---", id)
}
