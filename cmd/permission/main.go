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

func (dp fakeDataProvider) Get(ctx context.Context, fqfields ...definitions.Fqfield) ([]json.RawMessage, error) {
	m := make([]json.RawMessage, len(fqfields))
	for i := range fqfields {
		m[i] = json.RawMessage(strconv.Itoa(i))
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
