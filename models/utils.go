package models

func removeItem(items []*Item, itemToRemove *Item) []*Item {
	for i, item := range items {
		if item == itemToRemove {
			return append(items[:i], items[i+1:]...)
		}
	}
	return items
}
