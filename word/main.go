package main

import (
	"github.com/zhangdapeng520/zdpgo_cobra_examples/word/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd.Execute() error(%v)", err)
	}
}
