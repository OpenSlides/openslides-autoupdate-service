package main

import (
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"
	"github.com/OpenSlides/openslides-permission-service/pkg/permission"
)

type fakeDataProvider struct{}

func (dp fakeDataProvider) Get(fqfields []definitions.Fqfield) map[definitions.Fqfield]definitions.Value {
	m := make(map[definitions.Fqfield]definitions.Value)
	for i, val := range fqfields {
		m[val] = string(i)
	}
	return m
}

func main() {
	myDataProvider := fakeDataProvider{}
	ps := permission.New(myDataProvider)
	result, addition, err := ps.IsAllowed("", 0, nil)
	fmt.Println(result, addition, err)
	result, addition, err = ps.IsAllowed("topic.create", 0, nil)
	fmt.Println(result, addition, err)
	result, addition, err = ps.IsAllowed("topic.update", 0, nil)
	fmt.Println(result, addition, err)
	result, addition, err = ps.IsAllowed("topic.delete", 0, nil)
	fmt.Println(result, addition, err)
}
