package dataprovider

import "fmt"

type doesNotExistError string

func (e doesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", string(e))
}
