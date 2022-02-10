package example

type MyStruct struct {
	Name string
}

func (s *MyStruct) String() string {
	return s.Name
}

type Value interface {
	string | *MyStruct
}

type GenericMap[K string, V Value] map[K]V
