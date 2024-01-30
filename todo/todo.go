package todo

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 3 {
		return "(3)"
	}

	return " "
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

// ByPri implements sort.Interface for []Item based on
// priority and position.
type ByPri []Item

func (a ByPri) Len() int      { return len(a) }
func (a ByPri) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPri) Less(i, j int) bool {
	if a[i].Done != a[j].Done {
		return a[i].Done
	}

	if a[i].Priority == a[j].Priority {
		return a[i].position < a[j].position
	}

	return a[i].Priority < a[j].Priority
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "[X]"
	}

	return "[ ]"
}
