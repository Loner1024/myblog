package main

import (
	"github.com/Loner1024/uniix.io/configs"
	"log"
)

func main() {
	conf := configs.InitConfigs()
	s, err := wireApp(conf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(s.Start())
}
