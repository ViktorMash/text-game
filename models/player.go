package models

import (
	"fmt"
)

type Player struct {
	CurrentLocation *Room
	Items           []*Item
	HasBag          bool
}

func (p *Player) LookAround(location *Room) string {
	availableExits := location.GetExits()

	availableItems := "пустая комната"
	for i, f := range location.RoomFurniture {
		if len(f.Items) > 0 {
			comma := ", "
			if i == 0 {
				comma = ""
				availableItems = ""
			}
			availableItems = fmt.Sprintf("%s%sна %s: %s", availableItems, comma, f.NameInclined, location.GetItems(f))
		}
	}
	whereAreYou := ""
	if location.WhereAreYouNow != "" {
		whereAreYou = fmt.Sprintf("%s, ", location.WhereAreYouNow)
	}

	// получить задание для комнаты
	task := location.GetTask(p)
	
	return fmt.Sprintf("%s%s%s. можно пройти - %s", whereAreYou, availableItems, task, availableExits)
}

func (p *Player) GoTo(r *Room) string {
	// Проверяем, есть ли выход в эту комнату из текущей
	var targetExit *Exit
	for _, exit := range p.CurrentLocation.Exits {
		if exit.ToRoom == r {
			targetExit = exit
			break
		}
	}
	// Если выхода в комнату нет - отказ
	if targetExit == nil {
		return fmt.Sprintf("нет пути в %s", r.Name)
	}
	// Если выход найден и дверь закрыта
	if targetExit.IsDoorClosed {
		return "дверь закрыта"
	}

	availableExits := r.GetExits()
	statement := fmt.Sprintf("%s. можно пройти - %s", r.WhereYouCameFrom, availableExits)
	p.CurrentLocation = r
	return statement
}

func (p *Player) TakeItem(r *Room, itemName string) string {
	if itemName == "" {
		return "некорректные параметры"
	}
	itemIsBag := itemName == "рюкзак"
	playerHasBag := p.checkBag()

	// Вещи можно сложить только в рюкзак
	if !playerHasBag && !itemIsBag {
		return "некуда класть"
	}
	for _, f := range r.RoomFurniture {
		if len(f.Items) > 0 {
			for _, i := range f.Items {
				if i.Name == itemName {
					f.Items = removeItem(f.Items, i)
					p.Items = append(p.Items, i)

					if itemIsBag && !playerHasBag {
						p.HasBag = true
						return fmt.Sprintf("вы надели: %s", itemName)
					} else {
						return fmt.Sprintf("предмет добавлен в инвентарь: %s", itemName)
					}
				}
			}
		}
	}
	return "нет такого"
}

func (p *Player) ApplyItem(applyItem string, applyTo string) string {
	// Проверяем, есть ли предмет в инвентаре
	hasItem := false
	for _, i := range p.Items {
		if i.Name == applyItem {
			hasItem = true
			break
		}
	}

	if !hasItem {
		return fmt.Sprintf("нет предмета в инвентаре - %s", applyItem)
	}

	// Пробуем применить ключи к двери
	if applyItem == "ключи" && applyTo == "дверь" {
		// Ищем закрытую дверь в выходах из текущей комнаты
		for _, exit := range p.CurrentLocation.Exits {
			if exit.IsDoorClosed {
				exit.IsDoorClosed = false
				return "дверь открыта"
			}
		}
	}
	// описать другие юзкейсы для Items

	return "не к чему применить"
}

func (p *Player) checkBag() bool {
	if p.HasBag {
		return true
	}
	return false
}
