package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/TomiLabo/tmngparser"
)

// VERSION  version infomation
const VERSION = `0.1.0`

var (
	version bool
)

func init() {
	flag.BoolVar(&version, "v", false, "version")
}

func main() {
	flag.Parse()

	if version == true {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	var wg sync.WaitGroup
	for _, filename := range flag.Args() {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			data := parser.ReadFile(filename)
			ast := parser.Parse(data)
			json, err := json.Marshal(&ast)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(json))
		}(filename)
	}
	wg.Wait()
}
