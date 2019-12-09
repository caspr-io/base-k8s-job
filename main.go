package main

import "github.com/caspr-io/caspr-result/cmd"

import "github.com/caspr-io/mu-kit/kit"

func main() {
	kit.InitLogger("caspr-result")
	cmd.Execute()
}
