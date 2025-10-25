package main

import (
	"fmt"

	"github.com/ViktorMash/text-game/logic"
)

func main() {
	logic.InitGame()
	commands := []string{
		"осмотреться",
		"идти коридор",
		"идти комната",
		"надеть рюкзак",
		"взять ключи",
		"идти коридор",
		"применить ключи дверь",
		"идти улица",
	}

	for i, cmd := range commands {
		result := logic.HandleCommand(cmd)
		fmt.Printf("%d. %s: %s\n", i+1, cmd, result)
	}
}
