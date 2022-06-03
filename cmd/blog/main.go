package main

import (
	"github.com/Loner1024/uniix.io/configs"
	"github.com/Loner1024/uniix.io/logger"
	"log"
)

const serviceName = "Blog"

func main() {
	conf := configs.InitConfigs()
	l, err := logger.New("Blog")
	if err != nil {
		log.Fatalln(err)
	}
	s, err := wireApp(conf, l)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(s.Start())
}
