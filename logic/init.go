package logic

import (
	"github.com/ViktorMash/text-game/models"
)

func GameInit() (*models.Player, map[string]*models.Room) {
	rooms := make(map[string]*models.Room)

	// Кухня
	kitchenTableItems := []*models.Item{
		{Name: "чай"},
	}
	kitchenTable := &models.RoomFurniture{
		Name:         "стол",
		NameInclined: "столе",
		Items:        kitchenTableItems,
	}
	kitchen := &models.Room{
		Name:             "кухня",
		NameInclined:     "кухне",
		WhereYouCameFrom: "кухня, ничего интересного",
		WhereAreYouNow:   "ты находишься на кухне",
		RoomFurniture:    []*models.RoomFurniture{kitchenTable},
	}

	// комната
	roomTableItems := []*models.Item{
		{Name: "ключи", ApplyTo: []string{"дверь"}},
		{Name: "конспекты"},
	}
	roomTable := &models.RoomFurniture{
		Name:         "стол",
		NameInclined: "столе",
		Items:        roomTableItems,
	}
	roomChairItems := []*models.Item{
		{Name: "рюкзак"},
	}
	roomChair := &models.RoomFurniture{
		Name:         "стул",
		NameInclined: "стуле",
		Items:        roomChairItems,
	}
	room := &models.Room{
		Name:             "комната",
		NameInclined:     "комнате",
		WhereYouCameFrom: "ты в своей комнате",
		WhereAreYouNow:   "",
		RoomFurniture:    []*models.RoomFurniture{roomTable, roomChair},
	}

	// коридор
	corridor := &models.Room{
		Name:             "коридор",
		NameInclined:     "коридоре",
		WhereYouCameFrom: "ничего интересного",
		WhereAreYouNow:   "",
	}

	// улица
	street := &models.Room{
		Name:             "улица",
		NameInclined:     "улице",
		WhereYouCameFrom: "на улице весна",
		WhereAreYouNow:   "",
	}

	// инициализация выходов
	kitchen.Exits = []*models.Exit{{ToRoom: corridor}}
	room.Exits = []*models.Exit{{ToRoom: corridor}}
	corridor.Exits = []*models.Exit{
		{ToRoom: kitchen},
		{ToRoom: room},
		{ToRoom: street, IsDoorClosed: true},
	}

	// мапа с комнатами
	rooms[kitchen.Name] = kitchen
	rooms[corridor.Name] = corridor
	rooms[room.Name] = room
	rooms[street.Name] = street

	// игрок
	player := &models.Player{
		CurrentLocation: kitchen,
		Items:           []*models.Item{}, // ничего нет
		HasBag:          false,
	}

	return player, rooms
}
