package main

import (
	"errors"
	"fmt"
)

var FeesNotSubmitted = errors.New("fees not submitted")
var AdmissionCancelled = errors.New("admission cancelled")
var Foo = errors.New("foo")

func fees() error {

	return fmt.Errorf("%w", FeesNotSubmitted)
}

func admission() error {
	err := fees()
	if err != nil {
		return fmt.Errorf("%w %v", err, AdmissionCancelled)
	}
	return nil
}
func foo() error {
	err := admission()
	if err != nil {

		return fmt.Errorf("%w %v", err, Foo)
	}
	return nil
}

func main() {
	err := foo()
	fmt.Println(err)

	err = errors.Unwrap(err)

	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
}
