package main

import (
	"learn-go/dependency"
	"log"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	dependency.DemoV1(logger)
	dependency.DemoV2(logger.Println)
	dependency.DemoV3(logger)

}
