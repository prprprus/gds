package singlylinkedlist

import "encoding/json"

// ToJSON converts values of list to JSON.
func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

// FromJSON converts JSON to values of list.
func (list *List) FromJSON(data []byte) error {
	elements := make([]interface{}, 0)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		list.Clear()
		list.Append(elements...)
	}
	return err
}
