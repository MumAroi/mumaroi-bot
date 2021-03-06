package main

import (
	"flag"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	filename := flag.String("f", "bot.toml", "Set path for the config file.")
	flag.Parse()
	log.Println("INFO: Opening", *filename)
	err := LoadConfig(*filename)
	if err != nil {
		log.Println("FATA:", err)
		return
	}
}
