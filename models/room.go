package models

import (
	"fmt"
	"strings"
)

type Exit struct {
	ToRoom       *Room
	IsDoorClosed bool
}

type Room struct {
	Name             string
	NameInclined     string
	WhereYouCameFrom string
	WhereAreYouNow   string
	Exits            []*Exit
	RoomFurniture    []*RoomFurniture
}

func (r *Room) GetExits() string {
	availableExits := "нет выходов"
	if r.Name == "улица" {
		availableExits = "домой"
	}
	if len(r.Exits) > 0 {
		exits := make([]Named, len(r.Exits))
		for i, e := range r.Exits {
			exits[i] = e.ToRoom
		}
		availableExits = sliceToString(exits)
	}
	return availableExits
}

func (r *Room) GetItems(f *RoomFurniture) string {
	availableItems := fmt.Sprintf("на %s ничего нет", f.Name)
	if len(f.Items) > 0 {
		items := make([]Named, len(f.Items))
		for i, e := range f.Items {
			items[i] = e
		}
		availableItems = sliceToString(items)
	}
	return availableItems
}

func (r *Room) GetTask(player *Player) string {
	// задание при команде осмотреться на кухне
	if r.Name == "кухня" {
		gottaGetBag := ""
		if !player.HasBag {
			gottaGetBag = "собрать рюкзак и "
		}
		return fmt.Sprintf(", надо %sидти в универ", gottaGetBag)
	}
	return ""
}

// хотел обойтись одной функцией через интерфейс, минусы - создается еще один слайс
// Можно было написать 2 разные функции, минусы - дублирование кода
func (r *Room) GetName() string {
	return r.Name
}

type Named interface {
	GetName() string
}

func sliceToString(s []Named) string {
	stringsSlice := make([]string, len(s))
	for i, str := range s {
		stringsSlice[i] = str.GetName()
	}
	return strings.Join(stringsSlice, ", ")
}
