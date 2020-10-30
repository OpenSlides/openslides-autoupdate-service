package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
	"github.com/OpenSlides/openslides-permission-service/pkg/permission"
)

type fakeDataProvider struct{}

func (dp fakeDataProvider) Get(ctx context.Context, fqfields ...definitions.Fqfield) (map[definitions.Fqfield]json.RawMessage, error) {
	m := make(map[definitions.Fqfield]json.RawMessage)
	for i, val := range fqfields {
		m[val] = json.RawMessage(strconv.Itoa(i))
	}
	return m, nil
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
