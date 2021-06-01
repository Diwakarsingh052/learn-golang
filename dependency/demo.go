package dependency

import (
	"errors"
	"log"
)

// DemoV1 logs something on stdout
func DemoV1(log *log.Logger) { // strict exact dependency
	err := doThing()
	log.Println(err)
}

func DemoV2(logFn func(...interface{})) { // not true
	err := doThing()
	if err != nil {
		logFn(err)
	}
	log.Println()

}

type Logger interface {
	Println(...interface{})
	Printf(string, ...interface{})
}

// inject logger dependency and we don't care it is actual or fake
func DemoV3(log Logger) {
	err:= doThing()
	log.Println(err)
}

func doThing() error {
	return errors.New("error opening in file")
}
