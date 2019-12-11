package main

import (
	"github.com/caspr-io/caspr/cmd"
	"github.com/caspr-io/mu-kit/kit"
)

func main() {
	kit.InitLogger("caspr")
	cmd.Execute()
}
