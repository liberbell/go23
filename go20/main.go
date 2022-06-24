package main

import "log"

func main() {
	log.Panicln("logging!")
	log.Panicf("%T %v", "test", "test")
}
