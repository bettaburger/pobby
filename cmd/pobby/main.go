package main

import (
	"github.com/bettaburger/pobby/cmd"
	zone "github.com/lrstanley/bubblezone/v2"
)

func main() {
	zone.NewGlobal()
	cmd.Execute()
}