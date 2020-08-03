package main

import (
	"fmt"
)

type fakeDataProvider struct{}

type Fqfield = string
type Value = interface{}

func (dp *fakeDataProvider) get(fqfields []Fqfield) map[Fqfield]Value {
	m := make(map[Fqfield]Value)
	for i, val := range fqfields {
		m[val] = i
	}
	return m
}

func main() {
	myDataProvider := fakeDataProvider{}
	myFakeData := []Fqfield{"Foo", "Bar"}
	result := DoIt(myDataProvider, myFakeData)
	fmt.Println(result)
}
