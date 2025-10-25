package logic

import (
	"strings"
	"text-game/models"
)

var player *models.Player
var gameRooms map[string]*models.Room

func InitGame() {
	player, gameRooms = GameInit()
}

func HandleCommand(command string) string {
	parts := strings.Split(command, " ")
	action := parts[0]

	var param1, param2 string
	if len(parts) > 1 {
		param1 = parts[1]
	}
	if len(parts) > 2 {
		param2 = parts[2]
	}

	switch action {
	case "осмотреться":
		return player.LookAround(player.CurrentLocation)
	case "идти":
		return player.GoTo(gameRooms[param1]) // room
	case "надеть", "взять":
		return player.TakeItem(
			player.CurrentLocation, // room
			param1,                 // itemName
		)
	case "применить":
		if param1 != "" && param2 != "" {
			return player.ApplyItem(param1, param2)
		}
	}
	return "неизвестная команда"
}
