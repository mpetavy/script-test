package main

import (
	"embed"
	"github.com/mpetavy/common"
	"os"
	"time"
)

// install node module "hl7-standard" per "npm install -g hl7-standard" or local !DebugError(err)

//go:embed go.mod
var resources embed.FS

func init() {
	common.Init("", "", "", "", "test", "", "", "", &resources, nil, nil, run, 0)
}

func run() error {
	src, err := os.ReadFile("./test.js")
	if common.Error(err) {
		return err
	}

	se, err := common.NewScriptEngine(string(src), "")
	if common.Error(err) {
		return err
	}

	_, err = se.Run(time.Hour, "main", "")
	if common.Error(err) {
		return err
	}

	return nil
}

func main() {
	common.Run(nil)
}
