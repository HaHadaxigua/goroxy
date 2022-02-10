// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doublylinkedlist

import (
	"encoding/json"
)

// ToJSON outputs the JSON representation of list's elements.
func (list *List[V]) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

// FromJSON populates list's elements from the input JSON representation.
func (list *List[V]) FromJSON(data []byte) error {
	var elements []V
	err := json.Unmarshal(data, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}
