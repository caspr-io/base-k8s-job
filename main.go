package main

import (
	"github.com/caspr-io/caspr/cmd"
	mulog "github.com/caspr-io/mu-kit/log"
)

func main() {
	mulog.Init("caspr", &mulog.Config{Level: "info"})
	cmd.Execute()
}
