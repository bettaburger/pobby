package main

import (
	"github.com/bettaburger/pobby/internal/command"

	zone "github.com/lrstanley/bubblezone/v2"
)

func main() {
	zone.NewGlobal()
	command.Execute()
}