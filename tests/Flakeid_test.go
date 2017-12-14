package tests

import (
	"testing"
	"mhsyquanxian/quanju"
	"log"
)

func TestFlakeid(t *testing.T) {
	id, _ := quanju.Flakeid().NextId()
	log.Println("id---", id)
}
