package models

type Item struct {
	Name         string
	NameInclined string
	ApplyTo      []string
}

func (i *Item) GetName() string {
	return i.Name
}
