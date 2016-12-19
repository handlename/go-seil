package seil

import (
	"encoding/json"
	"os"

	apex "github.com/apex/go-apex"
)

// HandleFunc is a wrapper for apex.HandleFunc.
// In fn, output to stdout will be redirected to stderr
// to avoid parse error in apex using stdout as JSON.
func HandleFunc(fn func(event json.RawMessage, ctx *apex.Context) (res interface{}, err error)) {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		stdout := os.Stdout
		os.Stdout = os.Stderr
		defer func() { os.Stdout = stdout }()

		return fn(event, ctx)
	})
}
