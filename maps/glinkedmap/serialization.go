// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glinkedmap

import (
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v3"

	"github.com/HaHadaxigua/goroxy/utils"
	yamlUtil "github.com/ghodss/yaml"
)

// ToJSON outputs the JSON representation of map.
func (m *Map[K, V]) ToJSON() ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)

	buf.WriteRune('{')

	it := m.Iterator()
	lastIndex := m.Size() - 1
	index := 0

	for it.Next() {
		km, err := json.Marshal(it.Key())
		if err != nil {
			return nil, err
		}
		buf.Write(km)

		buf.WriteRune(':')

		vm, err := json.Marshal(it.Value())
		if err != nil {
			return nil, err
		}
		buf.Write(vm)

		if index != lastIndex {
			buf.WriteRune(',')
		}

		index++
	}

	buf.WriteRune('}')

	return buf.Bytes(), nil
}

// FromJSON populates map from the input JSON representation.
//func (m *Map) FromJSON(data []byte) error {
//	elements := make(map[string]interface{})
//	err := json.Unmarshal(data, &elements)
//	if err == nil {
//		m.Clear()
//		for key, value := range elements {
//			m.Put(key, value)
//		}
//	}
//	return err
//}

// FromJSON populates map from the input JSON representation.
func (m *Map[K, V]) FromJSON(data []byte) error {
	elements := make(map[K]V)
	err := json.Unmarshal(data, &elements)
	if err != nil {
		return err
	}

	index := make(map[K]int)
	var keys []K
	for key := range elements {
		keys = append(keys, key)
		esc, _ := json.Marshal(key)
		index[key] = bytes.Index(data, esc)
	}

	byIndex := func(a, b interface{}) int {
		key1 := a.(K)
		key2 := b.(K)
		index1 := index[key1]
		index2 := index[key2]
		return index1 - index2
	}

	utils.Sort(keys, byIndex)

	m.Clear()

	for _, key := range keys {
		m.Put(key, elements[key])
	}

	return nil
}

// -----------------------------------------------------------------------------

func (m *Map[K, V]) ToYaml() ([]byte, error) {
	buf, err := m.ToJSON()
	if err != nil {
		return nil, err
	}

	return yamlUtil.JSONToYAML(buf)
}

func (m *Map[K, V]) FromYaml(data []byte) error {
	jsonData, err := yamlUtil.YAMLToJSON(data) // convert yaml to json
	if err != nil {
		return err
	}

	if err = m.FromJSON(jsonData); err != nil {
		return err
	}
	return nil
}

func (m *Map[K, V]) ToYamlNode() (*yaml.Node, error) {
	var node yaml.Node
	buf, err := m.ToYaml()
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(buf, &node); err != nil {
		return nil, err
	}

	return &node, nil
}
