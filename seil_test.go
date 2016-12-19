package seil_test

import (
	"encoding/json"
	"fmt"

	apex "github.com/apex/go-apex"
	seil "github.com/handlename/go-seil"
)

type Result struct {
	Status  int
	Message string
}

func ExampleHandleFunc() {
	seil.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		fmt.Println("hello!!") // write to stderr.

		return Result{
			Status:  200,
			Message: "ok",
		}, nil
	})
}
