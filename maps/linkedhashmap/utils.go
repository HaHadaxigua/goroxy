package linkedhashmap

import "github.com/pkg/errors"

func (m *Map[K, V]) Merge(other *Map[K, V]) error {
	return other.AdvancedEach(func(key K, value V) error {
		if _, ok := m.Get(key); ok {
			return errors.Wrapf(ErrDuplicatedKey, "key: %v", key)
		}
		m.Put(key, value)
		return nil
	})
}

func (m *Map[K, V]) ToNormalMap() (map[K]V, error) {
	data := make(map[K]V)
	if err := m.AdvancedEach(func(key K, value V) error {
		data[key] = value
		return nil
	}); err != nil {
		return nil, err
	}

	return data, nil
}
