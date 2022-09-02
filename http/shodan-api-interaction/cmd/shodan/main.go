package main

import (
	"fmt"
	"log"
	"os"
	"shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Query Credits: %d\nScan Credis: %d\n\n", info.QueryCredits, info.ScanCredits)
	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	for _, host := range hostSearch.Matches {
		fmt.Println("%18s%8d\n", host.IPString, host.Port)
	}
}
