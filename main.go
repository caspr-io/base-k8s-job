package main

import "github.com/caspr-io/mu-kit/streaming"

import "github.com/caspr-io/mu-kit/kit"

import "context"

func main() {
	stanConfig := &streaming.Config{}
	if err := kit.ReadConfig("", stanConfig); err != nil {
		panic(err)
	}

	river, err := streaming.NewRiver(stanConfig)
	if err != nil {
		panic(err)
	}

	river.Publish(context.Background())
}
