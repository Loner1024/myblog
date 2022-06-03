package main

import (
	"log"
)

const serviceName = "Blog"

func main() {
	s, err := wireApp()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(s.Start())
}
