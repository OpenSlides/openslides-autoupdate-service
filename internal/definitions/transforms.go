package definitions

import (
	"fmt"
	"strconv"
	"strings"
)

var keyseparator = "/"

func CollectionAndIdFromFqid(fqid Fqid) (Collection, Id, error) {
	parts := strings.Split(fqid, keyseparator)
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("'%s' is not a valid fqid", fqid)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil || id <= 0 {
		return "", 0, fmt.Errorf("The id in '%s' is not valid", fqid)
	}

	return parts[0], id, nil
}

func FqidFromCollectionAndId(collection Collection, id Id) Fqid {
	return collection + keyseparator + strconv.Itoa(id)
}

func FqfieldFromFqidAndField(fqid Fqid, field Field) Fqid {
	return fqid + keyseparator + field
}
