package dataprovider

import "fmt"

// DoesNotExistError is thowen when an field does not exist.
type DoesNotExistError string

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf("%s does not exist.", string(e))
}
