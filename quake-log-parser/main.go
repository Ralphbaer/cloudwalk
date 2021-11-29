package main

import (
	"log"

	"github.com/Ralphbaer/cloudwalk/quake-log-parser/gen"
)

func main() {
	gen.InitializeApp().Run()
	log.Println("QuakeLogParser terminated")
}
