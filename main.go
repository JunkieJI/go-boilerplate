package main

import (
	"fmt"

	"github.com/JunkieJI/go-boilerplate/cmd"
	"github.com/JunkieJI/go-boilerplate/config"
	"github.com/JunkieJI/go-boilerplate/store"
)

func main() {
	cmd.Execute()

	config.Init()
	s := store.NewStore(config.GetDSN())
	fmt.Println(s.Ping())
}
