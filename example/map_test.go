package example

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	mem := make(GenericMap[string, *MyStruct])

	mem["1"] = &MyStruct{
		Name: "1",
	}

	for k, v := range mem {
		fmt.Printf("key: %s \n", k)
		fmt.Printf("value: %s \n", v.String())
	}
}
