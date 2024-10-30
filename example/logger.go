package main

import "log"

func customLog() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("Testing log format")
}
